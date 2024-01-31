package common

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// UISysSimSpeed is a simple system that changes the simulation speed
// when the user presses PageUp, PageDown or Home.
type UISysSimSpeed struct {
	MinExponent int
	MaxExponent int

	speed generic.Resource[SimulationSpeed]
}

// InitializeUI the system
func (s *UISysSimSpeed) InitializeUI(world *ecs.World) {
	s.speed = generic.NewResource[SimulationSpeed](world)
	if !s.speed.Has() {
		ecs.AddResource(world, &SimulationSpeed{})
	}
}

// UpdateUI the system
func (s *UISysSimSpeed) UpdateUI(world *ecs.World) {
	if inpututil.IsKeyJustPressed(ebiten.KeyPageUp) {
		speed := s.speed.Get()
		if speed.Exponent < s.MaxExponent {
			speed.Exponent++
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyPageDown) {
		speed := s.speed.Get()
		if speed.Exponent > s.MinExponent {
			speed.Exponent--
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyHome) {
		speed := s.speed.Get()
		speed.Exponent = 0
	}
}

// PostUpdateUI the system
func (s *UISysSimSpeed) PostUpdateUI(world *ecs.World) {}

// FinalizeUI the system
func (s *UISysSimSpeed) FinalizeUI(world *ecs.World) {}
