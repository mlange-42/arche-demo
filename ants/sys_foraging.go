package ants

import (
	"math/rand"

	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// SysForaging is a system that performs forager decisions.
type SysForaging struct {
	MaxCollect float64
	RandomProb float64

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
}

// Initialize the system
func (s *SysForaging) Initialize(world *ecs.World) {
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
	query := s.filter.Query(world)
	for query.Next() {
		antEdge, _ := query.Get()
		if antEdge.Pos < antEdge.Length {
			continue
		}
		oldEdge, oldTrace := s.edgeMap.Get(antEdge.Edge)
		oldTrace.FromNest++

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

	maxTrace := -1.0
	var maxEdge ecs.Entity
	count := 0
	for i := 0; i < node.NumEdges; i++ {
		edge := node.InEdges[i]
		trace := s.traceMap.Get(edge)
		if trace.FromResource > maxTrace {
			maxTrace = trace.FromResource
			maxEdge = node.OutEdges[i]
			count = 1
			continue
		}
		if trace.FromResource == maxTrace {
			count++
			if rand.Float64() < 1.0/float64(count) {
				maxEdge = edge
			}
		}
	}
	return maxEdge
}
