package ants

import (
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// SysMoveAnts is a system that moves ants along their edges.
type SysMoveAnts struct {
	MaxSpeed float64

	filter generic.Filter2[AntEdge, Position]
}

// Initialize the system
func (s *SysMoveAnts) Initialize(world *ecs.World) {
	s.filter = *generic.NewFilter2[AntEdge, Position]()
}

// Update the system
func (s *SysMoveAnts) Update(world *ecs.World) {
	query := s.filter.Query(world)
	for query.Next() {
		antEdge, pos := query.Get()
		antEdge.Pos += s.MaxSpeed
		antEdge.UpdatePos(pos)
	}
}

// Finalize the system
func (s *SysMoveAnts) Finalize(world *ecs.World) {}
