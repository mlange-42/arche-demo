package main

import (
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// Box2DPhysics system
type Box2DPhysics struct {
	worldRes generic.Resource[BoxWorld]
}

// Initialize the system
func (s *Box2DPhysics) Initialize(world *ecs.World) {
	s.worldRes = generic.NewResource[BoxWorld](world)
}

// Update the system
func (s *Box2DPhysics) Update(world *ecs.World) {
	w := s.worldRes.Get().World

	w.Step(1.0/60.0, 5, 2)
}

// Finalize the system
func (s *Box2DPhysics) Finalize(world *ecs.World) {}
