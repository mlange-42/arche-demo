package ants

import (
	"math"
	"math/rand"

	"github.com/mlange-42/arche-demo/common"
	"github.com/mlange-42/arche-model/resource"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// SysReturning is a system that performs decisions of ants going back to their nest.
type SysReturning struct {
	ProbExponent float64
	RandomProb   float64
	TraceDecay   float64

	time        generic.Resource[resource.Tick]
	nest        generic.Resource[Nest]
	filter      generic.Filter2[AntEdge, ActReturn]
	antEdgeMap  generic.Map1[AntEdge]
	edgeMap     generic.Map2[Edge, Trace]
	edgeGeomMap generic.Map1[EdgeGeometry]
	nodeMap     generic.Map1[Node]
	traceMap    generic.Map1[Trace]

	exchangeArrive generic.Exchange

	toArrive []ecs.Entity
	probs    [8]float64
}

// Initialize the system
func (s *SysReturning) Initialize(world *ecs.World) {
	s.time = generic.NewResource[resource.Tick](world)
	s.nest = generic.NewResource[Nest](world)
	s.filter = *generic.NewFilter2[AntEdge, ActReturn]()

	s.antEdgeMap = generic.NewMap1[AntEdge](world)
	s.edgeMap = generic.NewMap2[Edge, Trace](world)
	s.edgeGeomMap = generic.NewMap1[EdgeGeometry](world)
	s.nodeMap = generic.NewMap1[Node](world)
	s.traceMap = generic.NewMap1[Trace](world)

	s.exchangeArrive = *generic.NewExchange(world).
		Adds(generic.T[ActInNest]()).
		Removes(generic.T2[ActReturn, AntEdge]()...)

	s.toArrive = make([]ecs.Entity, 0, 16)
}

// Update the system
func (s *SysReturning) Update(world *ecs.World) {
	tick := s.time.Get().Tick
	nest := s.nest.Get()

	query := s.filter.Query(world)
	for query.Next() {
		antEdge, ret := query.Get()
		if antEdge.Pos < antEdge.Length {
			continue
		}
		oldEdge, oldTrace := s.edgeMap.Get(antEdge.Edge)
		oldTrace.FromResource += ret.Load * math.Pow(s.TraceDecay, float64(tick-ret.Start))

		if oldEdge.To == nest.Node {
			s.toArrive = append(s.toArrive, query.Entity())
			continue
		}
		node := s.nodeMap.Get(oldEdge.To)
		edge := s.selectEdge(world, node)
		geom := s.edgeGeomMap.Get(edge)

		antEdge.Update(edge, geom)
	}

	for _, e := range s.toArrive {
		s.exchangeArrive.Exchange(e)
	}

	s.toArrive = s.toArrive[:0]
}

// Finalize the system
func (s *SysReturning) Finalize(world *ecs.World) {}

func (s *SysReturning) selectEdge(world *ecs.World, node *Node) ecs.Entity {
	if rand.Float64() < s.RandomProb {
		return node.OutEdges[rand.Intn(node.NumEdges)]
	}

	for i := 0; i < node.NumEdges; i++ {
		edge := node.InEdges[i]
		trace := s.traceMap.Get(edge)
		s.probs[i] = math.Pow(trace.FromNest, s.ProbExponent)
	}

	if sel, ok := common.SelectRoulette(s.probs[:node.NumEdges]); ok {
		return node.OutEdges[sel]
	}
	return node.OutEdges[rand.Intn(node.NumEdges)]
}
