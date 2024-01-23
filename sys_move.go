package main

import (
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// MoveEntities system
type MoveEntities struct {
	filter generic.Filter2[Position, Target]
}

// Initialize the system
func (s *MoveEntities) Initialize(world *ecs.World) {
	s.filter = *generic.NewFilter2[Position, Target]()
}

// Update the system
func (s *MoveEntities) Update(world *ecs.World) {
	query := s.filter.Query(world)
	for query.Next() {
		pos, targ := query.Get()
		dx := targ.X - pos.X
		dy := targ.Y - pos.Y
		pos.X += 0.01 * dx
		pos.Y += 0.01 * dy
	}
}

// Finalize the system
func (s *MoveEntities) Finalize(world *ecs.World) {}
