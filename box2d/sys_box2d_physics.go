package box2d

import (
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// SysB2Physics is a simple system that updates the Box2D world to perform a physics step.
type SysB2Physics struct {
	worldRes generic.Resource[BoxWorld]
}

// Initialize the system
func (s *SysB2Physics) Initialize(world *ecs.World) {
	s.worldRes = generic.NewResource[BoxWorld](world)
}

// Update the system
func (s *SysB2Physics) Update(world *ecs.World) {
	w := s.worldRes.Get().World

	w.Step(1.0/60.0, 5, 2)
}

// Finalize the system
func (s *SysB2Physics) Finalize(world *ecs.World) {}
