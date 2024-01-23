package main

import (
	"math"
	"math/rand"

	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// MoveEntities system
type MoveEntities struct {
	MaxSpeed     float64
	MaxAcc       float64
	FleeDistance float64
	Damp         float64
	canvas       generic.Resource[Canvas]
	filter       generic.Filter3[Position, Velocity, Target]
	counter      uint64
}

// Initialize the system
func (s *MoveEntities) Initialize(world *ecs.World) {
	s.canvas = generic.NewResource[Canvas](world)
	s.filter = *generic.NewFilter3[Position, Velocity, Target]()
}

// Update the system
func (s *MoveEntities) Update(world *ecs.World) {
	mouse := s.canvas.Get().Mouse
	_ = mouse

	query := s.filter.Query(world)
	for query.Next() {
		pos, vel, targ := query.Get()

		attrX, attrY, _ := s.norm(targ.X-pos.X, targ.Y-pos.Y)

		vel.X += attrX * s.MaxAcc
		vel.Y += attrY * s.MaxAcc

		velAbs := math.Sqrt(vel.X*vel.X + vel.Y*vel.Y)
		if velAbs > 1.0 {
			vel.X /= velAbs
			vel.Y /= velAbs
			velAbs = 1.0
		}
		if s.counter%23 == 0 {
			vel.X += rand.NormFloat64() * velAbs * 0.1
			vel.Y += rand.NormFloat64() * velAbs * 0.1
		}

		vel.X *= s.Damp
		vel.Y *= s.Damp

		pos.X += vel.X * s.MaxSpeed
		pos.Y += vel.Y * s.MaxSpeed

		s.counter++
	}
}

func (s *MoveEntities) norm(dx, dy float64) (float64, float64, float64) {
	len := math.Sqrt(dx*dx + dy*dy)
	if len == 0 {
		return 0, 0, 0
	}
	return dx / len, dy / len, len
}

// Finalize the system
func (s *MoveEntities) Finalize(world *ecs.World) {}
