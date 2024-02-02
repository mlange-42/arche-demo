package systems

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/mlange-42/arche-demo/common"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// SimSpeed is a simple system that changes the simulation speed
// when the user presses PageUp, PageDown or Home.
type SimSpeed struct {
	InitialExponent int
	MinExponent     int
	MaxExponent     int

	speed generic.Resource[common.SimulationSpeed]
}

// InitializeUI the system
func (s *SimSpeed) InitializeUI(world *ecs.World) {
	if s.MinExponent > s.MaxExponent {
		panic("min exponent must not be higher than max exponent")
	}
	if s.InitialExponent < s.MinExponent || s.InitialExponent > s.MaxExponent {
		panic("initial exponent must be in range min/max exponent")
	}

	s.speed = generic.NewResource[common.SimulationSpeed](world)
	if s.speed.Has() {
		s.speed.Get().Exponent = s.InitialExponent
	} else {
		ecs.AddResource(world, &common.SimulationSpeed{Exponent: s.InitialExponent})
	}
}

// UpdateUI the system
func (s *SimSpeed) UpdateUI(world *ecs.World) {
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
func (s *SimSpeed) PostUpdateUI(world *ecs.World) {}

// FinalizeUI the system
func (s *SimSpeed) FinalizeUI(world *ecs.World) {}
