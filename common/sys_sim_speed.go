package common

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// SysSimSpeed is a simple system that changes the simulation speed
// when the user presses PageUp, PageDown or Home.
type SysSimSpeed struct {
	InitialExponent int
	MinExponent     int
	MaxExponent     int

	speed generic.Resource[SimulationSpeed]
}

// Initialize the system
func (s *SysSimSpeed) Initialize(world *ecs.World) {
	if s.MinExponent > s.MaxExponent {
		panic("min exponent must not be higher than max exponent")
	}
	if s.InitialExponent < s.MinExponent || s.InitialExponent > s.MaxExponent {
		panic("initial exponent must be in range min/max exponent")
	}

	s.speed = generic.NewResource[SimulationSpeed](world)
	if s.speed.Has() {
		s.speed.Get().Exponent = s.InitialExponent
	} else {
		ecs.AddResource(world, &SimulationSpeed{Exponent: s.InitialExponent})
	}
}

// Update the system
func (s *SysSimSpeed) Update(world *ecs.World) {
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

// Finalize the system
func (s *SysSimSpeed) Finalize(world *ecs.World) {}
