package ants

import (
	"math"
	"math/rand"

	"github.com/mlange-42/arche-demo/common"
	"github.com/mlange-42/arche-model/resource"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// SysForaging is a system that performs forager decisions.
type SysForaging struct {
	MaxCollect   float64
	ProbExponent float64
	RandomProb   float64
	TraceDecay   float64

	time        generic.Resource[resource.Tick]
	filter      generic.Filter2[AntEdge, ActForage]
	antEdgeMap  generic.Map1[AntEdge]
	returnMap   generic.Map1[ActReturn]
	edgeMap     generic.Map2[Edge, Trace]
	edgeGeomMap generic.Map1[EdgeGeometry]
	nodeMap     generic.Map1[Node]
	resourceMap generic.Map[NodeResource]
	traceMap    generic.Map1[Trace]

	exchangeReturn generic.Exchange

	toReturn []returnEntry
	probs    [8]float64
}

// Initialize the system
func (s *SysForaging) Initialize(world *ecs.World) {
	s.time = generic.NewResource[resource.Tick](world)
	s.filter = *generic.NewFilter2[AntEdge, ActForage]()

	s.antEdgeMap = generic.NewMap1[AntEdge](world)
	s.returnMap = generic.NewMap1[ActReturn](world)
	s.edgeMap = generic.NewMap2[Edge, Trace](world)
	s.edgeGeomMap = generic.NewMap1[EdgeGeometry](world)
	s.nodeMap = generic.NewMap1[Node](world)
	s.resourceMap = generic.NewMap[NodeResource](world)
	s.traceMap = generic.NewMap1[Trace](world)

	s.exchangeReturn = *generic.NewExchange(world).
		Adds(generic.T[ActReturn]()).
		Removes(generic.T[ActForage]())

	s.toReturn = make([]returnEntry, 0, 16)
}

// Update the system
func (s *SysForaging) Update(world *ecs.World) {
	tick := s.time.Get().Tick

	query := s.filter.Query(world)
	for query.Next() {
		antEdge, forage := query.Get()
		if antEdge.Pos < antEdge.Length {
			continue
		}
		oldEdge, oldTrace := s.edgeMap.Get(antEdge.Edge)
		oldTrace.FromNest += math.Pow(s.TraceDecay, float64(tick-forage.Start))

		if !s.resourceMap.Has(oldEdge.To) {
			node := s.nodeMap.Get(oldEdge.To)
			edge := s.selectEdge(world, node)
			geom := s.edgeGeomMap.Get(edge)

			antEdge.Update(edge, geom)
			continue
		}
		res := s.resourceMap.Get(oldEdge.To)

		entry := returnEntry{Entity: query.Entity(), Load: res.Resource}
		res.Resource -= s.MaxCollect
		if res.Resource < 0 {
			res.Resource = 0
		}
		s.toReturn = append(s.toReturn, entry)
	}

	for _, e := range s.toReturn {
		s.exchangeReturn.Exchange(e.Entity)
		ret := s.returnMap.Get(e.Entity)
		ret.Start = tick
		ret.Load = e.Load
	}

	s.toReturn = s.toReturn[:0]
}

// Finalize the system
func (s *SysForaging) Finalize(world *ecs.World) {}

func (s *SysForaging) selectEdge(world *ecs.World, node *Node) ecs.Entity {
	if rand.Float64() < s.RandomProb {
		return node.OutEdges[rand.Intn(node.NumEdges)]
	}

	for i := 0; i < node.NumEdges; i++ {
		edge := node.InEdges[i]
		trace := s.traceMap.Get(edge)
		s.probs[i] = math.Pow(trace.FromResource, s.ProbExponent)
	}

	if sel, ok := common.SelectRoulette(s.probs[:node.NumEdges]); ok {
		return node.OutEdges[sel]
	}
	return node.OutEdges[rand.Intn(node.NumEdges)]
}
