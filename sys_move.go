package main

import (
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// MoveEntities system
type MoveEntities struct {
	filter generic.Filter1[Position]
}

// Initialize the system
func (s *MoveEntities) Initialize(world *ecs.World) {
	s.filter = *generic.NewFilter1[Position]()
}

// Update the system
func (s *MoveEntities) Update(world *ecs.World) {
	query := s.filter.Query(world)
	for query.Next() {
		pos := query.Get()
		pos.X++
	}
}

// Finalize the system
func (s *MoveEntities) Finalize(world *ecs.World) {}
