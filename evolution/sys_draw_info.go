package evolution

import (
	"fmt"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/mlange-42/arche-demo/common"
	"github.com/mlange-42/arche-model/resource"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// UISysDrawInfo is a system that draws info text.
type UISysDrawInfo struct {
	Offset common.Vec2i

	canvas generic.Resource[common.EbitenImage]
	time   generic.Resource[resource.Tick]
	speed  generic.Resource[common.SimulationSpeed]

	filter generic.Filter0
}

// InitializeUI the system
func (s *UISysDrawInfo) InitializeUI(world *ecs.World) {
	s.canvas = generic.NewResource[common.EbitenImage](world)
	s.time = generic.NewResource[resource.Tick](world)
	s.speed = generic.NewResource[common.SimulationSpeed](world)

	s.filter = *generic.NewFilter0().With(generic.T[Age]())
}

// UpdateUI the system
func (s *UISysDrawInfo) UpdateUI(world *ecs.World) {
	tick := s.time.Get().Tick
	canvas := s.canvas.Get()
	screen := canvas.Image

	speed := math.Pow(2, float64(s.speed.Get().Exponent))

	query := s.filter.Query(world)
	entities := query.Count()
	query.Close()

	text := fmt.Sprintf(`FPS %.1f
Tick
  %d
Speed
  x %.2f
Entities
  %d`, ebiten.ActualFPS(), tick, speed, entities)

	ebitenutil.DebugPrintAt(screen, text, s.Offset.X, s.Offset.Y)
}

// PostUpdateUI the system
func (s *UISysDrawInfo) PostUpdateUI(world *ecs.World) {}

// FinalizeUI the system
func (s *UISysDrawInfo) FinalizeUI(world *ecs.World) {}
