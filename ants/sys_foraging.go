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
	MaxCollect       float64
	ProbExponent     float64
	RandomProb       float64
	TraceDecay       float64
	MaxSearchTime    int
	ScoutProbability float64

	time        generic.Resource[resource.Tick]
	filter      generic.Filter3[AntEdge, ActForage, Position]
	antEdgeMap  generic.Map1[AntEdge]
	returnMap   generic.Map1[ActReturn]
	scoutMap    generic.Map1[ActScout]
	forageMap   generic.Map1[ActForage]
	edgeMap     generic.Map2[Edge, Trace]
	edgeGeomMap generic.Map1[EdgeGeometry]
	nodeMap     generic.Map1[Node]
	resourceMap generic.Map[NodeResource]
	traceMap    generic.Map1[Trace]

	exchangeReturn generic.Exchange
	exchangeScout  generic.Exchange

	toReturn []returnEntry
	toScout  []ecs.Entity
	probs    [8]float64
}

// Initialize the system
func (s *SysForaging) Initialize(world *ecs.World) {
	s.time = generic.NewResource[resource.Tick](world)
	s.filter = *generic.NewFilter3[AntEdge, ActForage, Position]()

	s.antEdgeMap = generic.NewMap1[AntEdge](world)
	s.returnMap = generic.NewMap1[ActReturn](world)
	s.scoutMap = generic.NewMap1[ActScout](world)
	s.forageMap = generic.NewMap1[ActForage](world)
	s.edgeMap = generic.NewMap2[Edge, Trace](world)
	s.edgeGeomMap = generic.NewMap1[EdgeGeometry](world)
	s.nodeMap = generic.NewMap1[Node](world)
	s.resourceMap = generic.NewMap[NodeResource](world)
	s.traceMap = generic.NewMap1[Trace](world)

	s.exchangeReturn = *generic.NewExchange(world).
		Adds(generic.T[ActReturn]()).
		Removes(generic.T[ActForage]())
	s.exchangeScout = *generic.NewExchange(world).
		Adds(generic.T[ActScout]()).
		Removes(generic.T[ActForage]())

	s.toReturn = make([]returnEntry, 0, 16)
	s.toScout = make([]ecs.Entity, 0, 16)
}

// Update the system
func (s *SysForaging) Update(world *ecs.World) {
	tick := s.time.Get().Tick

	query := s.filter.Query(world)
	for query.Next() {
		antEdge, forage, pos := query.Get()
		newPos := antEdge.Pos - antEdge.Length
		if newPos < 0 {
			continue
		}
		oldEdge, oldTrace := s.edgeMap.Get(antEdge.Edge)
		oldTrace.FromNest += math.Pow(s.TraceDecay, float64(tick-forage.Start))

		if s.resourceMap.Has(oldEdge.To) {
			res := s.resourceMap.Get(oldEdge.To)

			entry := returnEntry{Entity: query.Entity(), Load: res.Resource}
			res.Resource -= s.MaxCollect
			if res.Resource < 0 {
				res.Resource = 0
			}
			s.toReturn = append(s.toReturn, entry)
			continue
		}
		if tick > forage.Start+int64(s.MaxSearchTime) {
			if rand.Float64() < s.ScoutProbability {
				s.toScout = append(s.toScout, query.Entity())
			} else {
				s.toReturn = append(s.toReturn, returnEntry{Entity: query.Entity(), Load: 0})
			}
			continue
		}

		node := s.nodeMap.Get(oldEdge.To)
		edge := s.selectEdge(world, node)
		geom := s.edgeGeomMap.Get(edge)

		antEdge.Update(edge, geom)
		antEdge.Pos = newPos
		antEdge.UpdatePos(pos)
	}

	for _, e := range s.toScout {
		start := s.forageMap.Get(e).Start
		s.exchangeScout.Exchange(e)
		scout := s.scoutMap.Get(e)
		scout.Start = start
	}

	for _, e := range s.toReturn {
		s.exchangeReturn.Exchange(e.Entity)
		ret := s.returnMap.Get(e.Entity)
		ret.Start = tick
		ret.Load = e.Load
	}

	s.toScout = s.toScout[:0]
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
