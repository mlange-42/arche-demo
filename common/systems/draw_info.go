package systems

import (
	"fmt"
	"image"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/mlange-42/arche-demo/common"
	"github.com/mlange-42/arche-model/resource"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// DrawInfo is a system that draws info text.
type DrawInfo struct {
	// Where to put the text.
	Offset image.Point
	// Optionally, components required for a second entity count entry.
	Components []generic.Comp

	canvas generic.Resource[common.EbitenImage]
	time   generic.Resource[resource.Tick]
	speed  generic.Resource[common.SimulationSpeed]

	filterAll generic.Filter0
	filter    generic.Filter0
}

// InitializeUI the system
func (s *DrawInfo) InitializeUI(world *ecs.World) {
	s.canvas = generic.NewResource[common.EbitenImage](world)
	s.time = generic.NewResource[resource.Tick](world)
	s.speed = generic.NewResource[common.SimulationSpeed](world)

	s.filterAll = *generic.NewFilter0()
	s.filter = *generic.NewFilter0().With(s.Components...)
}

// UpdateUI the system
func (s *DrawInfo) UpdateUI(world *ecs.World) {
	tick := s.time.Get().Tick
	canvas := s.canvas.Get()
	screen := canvas.Image

	speed := 1.0
	if s.speed.Has() {
		speed = math.Pow(2, float64(s.speed.Get().Exponent))
	}

	query := s.filterAll.Query(world)
	entities := query.Count()
	query.Close()

	text := fmt.Sprintf(`FPS %.1f
Tick
  %d
Speed
  x %.2f
Entities
  %d`, ebiten.ActualFPS(), tick, speed, entities)

	if len(s.Components) > 0 {
		query := s.filter.Query(world)
		entities := query.Count()
		query.Close()
		text += fmt.Sprintf("\n  (%d)", entities)
	}

	ebitenutil.DebugPrintAt(screen, text, s.Offset.X, s.Offset.Y)
}

// PostUpdateUI the system
func (s *DrawInfo) PostUpdateUI(world *ecs.World) {}

// FinalizeUI the system
func (s *DrawInfo) FinalizeUI(world *ecs.World) {}
