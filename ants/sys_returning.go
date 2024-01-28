package ants

import (
	"math/rand"

	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// SysReturning is a system that performs decisions of ants going back to their nest.
type SysReturning struct {
	RandomProb float64

	nest        generic.Resource[Nest]
	filter      generic.Filter2[AntEdge, ActReturn]
	antEdgeMap  generic.Map1[AntEdge]
	edgeMap     generic.Map2[Edge, Trace]
	edgeGeomMap generic.Map1[EdgeGeometry]
	nodeMap     generic.Map1[Node]
	traceMap    generic.Map1[Trace]

	exchangeArrive generic.Exchange

	toArrive []ecs.Entity
}

// Initialize the system
func (s *SysReturning) Initialize(world *ecs.World) {
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
	nest := s.nest.Get()

	query := s.filter.Query(world)
	for query.Next() {
		antEdge, ret := query.Get()
		if antEdge.Pos < antEdge.Length {
			continue
		}
		oldEdge, oldTrace := s.edgeMap.Get(antEdge.Edge)
		oldTrace.FromResource += ret.Load

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

	maxTrace := -1.0
	var maxEdge ecs.Entity
	count := 0
	for i := 0; i < node.NumEdges; i++ {
		edge := node.InEdges[i]
		trace := s.traceMap.Get(edge)
		if trace.FromNest > maxTrace {
			maxTrace = trace.FromNest
			maxEdge = node.OutEdges[i]
			count = 1
			continue
		}
		if trace.FromNest == maxTrace {
			count++
			if rand.Float64() < 1.0/float64(count) {
				maxEdge = edge
			}
		}
	}
	return maxEdge
}
