package ants

import (
	"math/rand"

	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

type returnEntry struct {
	Entity ecs.Entity
	Load   float64
}

// SysScouting is a system that performs scout decisions.
type SysScouting struct {
	MaxCollect float64

	filter      generic.Filter2[AntEdge, ActScout]
	antEdgeMap  generic.Map1[AntEdge]
	returnMap   generic.Map1[ActReturn]
	edgeMap     generic.Map2[Edge, Trace]
	edgeGeomMap generic.Map1[EdgeGeometry]
	nodeMap     generic.Map1[Node]
	resourceMap generic.Map[NodeResource]

	exchangeReturn generic.Exchange

	toReturn []returnEntry
}

// Initialize the system
func (s *SysScouting) Initialize(world *ecs.World) {
	s.filter = *generic.NewFilter2[AntEdge, ActScout]()

	s.antEdgeMap = generic.NewMap1[AntEdge](world)
	s.returnMap = generic.NewMap1[ActReturn](world)
	s.edgeMap = generic.NewMap2[Edge, Trace](world)
	s.edgeGeomMap = generic.NewMap1[EdgeGeometry](world)
	s.nodeMap = generic.NewMap1[Node](world)
	s.resourceMap = generic.NewMap[NodeResource](world)

	s.exchangeReturn = *generic.NewExchange(world).
		Adds(generic.T[ActReturn]()).
		Removes(generic.T[ActScout]())

	s.toReturn = make([]returnEntry, 0, 16)
}

// Update the system
func (s *SysScouting) Update(world *ecs.World) {
	query := s.filter.Query(world)
	for query.Next() {
		antEdge, _ := query.Get()
		if antEdge.Pos < antEdge.Length {
			continue
		}
		oldEdge, oldTrace := s.edgeMap.Get(antEdge.Edge)
		oldTrace.FromNest++

		node := s.nodeMap.Get(oldEdge.To)
		if !s.resourceMap.Has(oldEdge.To) {
			edge := node.OutEdges[rand.Intn(node.NumEdges)]
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
func (s *SysScouting) Finalize(world *ecs.World) {}
