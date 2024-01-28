package ants

import (
	"math/rand"

	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// SysScouting is a system that performs scout decisions.
type SysScouting struct {
	filter      generic.Filter2[AntEdge, ActScout]
	antEdgeMap  generic.Map1[AntEdge]
	edgeMap     generic.Map2[Edge, Trace]
	edgeGeomMap generic.Map1[EdgeGeometry]
	nodeMap     generic.Map1[Node]
}

// Initialize the system
func (s *SysScouting) Initialize(world *ecs.World) {
	s.filter = *generic.NewFilter2[AntEdge, ActScout]()

	s.antEdgeMap = generic.NewMap1[AntEdge](world)
	s.edgeMap = generic.NewMap2[Edge, Trace](world)
	s.edgeGeomMap = generic.NewMap1[EdgeGeometry](world)
	s.nodeMap = generic.NewMap1[Node](world)
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

		edge := node.OutEdges[rand.Intn(node.NumEdges)]
		geom := s.edgeGeomMap.Get(edge)

		antEdge.Update(edge, geom)
	}
}

// Finalize the system
func (s *SysScouting) Finalize(world *ecs.World) {}
