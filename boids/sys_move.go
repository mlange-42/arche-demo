package boids

import (
	"github.com/mlange-42/arche-demo/common"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// MoveEntities is a system that moves entities around.
type MoveEntities struct {
	canvas generic.Resource[common.Image]
	filter generic.Filter2[Position, Velocity]
}

// Initialize the system
func (s *MoveEntities) Initialize(world *ecs.World) {
	s.canvas = generic.NewResource[common.Image](world)
	s.filter = *generic.NewFilter2[Position, Velocity]()
}

// Update the system
func (s *MoveEntities) Update(world *ecs.World) {
	canvas := s.canvas.Get()
	w := float64(canvas.Width)
	h := float64(canvas.Height)

	query := s.filter.Query(world)
	for query.Next() {
		pos, vel := query.Get()

		pos.X += vel.X
		pos.Y += vel.Y

		if pos.X < 0 || pos.X >= w {
			vel.X *= -1
			pos.X += vel.X * 2
		}
		if pos.Y < 0 || pos.Y >= h {
			vel.Y *= -1
			pos.Y += vel.Y * 2
		}
	}
}

// Finalize the system
func (s *MoveEntities) Finalize(world *ecs.World) {}
