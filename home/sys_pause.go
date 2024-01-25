package main

import (
	"github.com/mlange-42/arche-model/model"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// ManagePause system
type ManagePause struct {
	systems generic.Resource[model.Systems]
	mouse   generic.Resource[MouseListener]
}

// InitializeUI the system
func (s *ManagePause) InitializeUI(world *ecs.World) {
	s.systems = generic.NewResource[model.Systems](world)
	s.mouse = generic.NewResource[MouseListener](world)
}

// UpdateUI the system
func (s *ManagePause) UpdateUI(world *ecs.World) {
	s.systems.Get().Paused = s.mouse.Get().Paused
}

// PostUpdateUI the system
func (s *ManagePause) PostUpdateUI(world *ecs.World) {}

// FinalizeUI the system
func (s *ManagePause) FinalizeUI(world *ecs.World) {}
