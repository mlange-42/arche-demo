package main

import (
	"math"

	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// MoveEntities system
type MoveEntities struct {
	MaxSpeed float64
	filter   generic.Filter2[Position, Target]
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
		len := math.Sqrt(dx*dx + dy*dy)
		if len < 0.05 {
			continue
		}
		scale := 1.0
		if len > s.MaxSpeed {
			scale = 1.0 / len
		}
		pos.X += dx * scale
		pos.Y += dy * scale
	}
}

// Finalize the system
func (s *MoveEntities) Finalize(world *ecs.World) {}
