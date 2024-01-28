package ants

import (
	"image"
	"image/color"
	"image/draw"

	"github.com/mlange-42/arche-demo/common"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// DrawResources is a system that draws resource patches.
type DrawResources struct {
	canvas     generic.Resource[common.Image]
	nodeFilter generic.Filter2[Position, NodeResource]
}

// InitializeUI the system
func (s *DrawResources) InitializeUI(world *ecs.World) {
	s.canvas = generic.NewResource[common.Image](world)
	s.nodeFilter = *generic.NewFilter2[Position, NodeResource]()
}

// UpdateUI the system
func (s *DrawResources) UpdateUI(world *ecs.World) {
	canvas := s.canvas.Get()
	img := canvas.Image

	rad := 5.0

	nodeQuery := s.nodeFilter.Query(world)
	for nodeQuery.Next() {
		pos, res := nodeQuery.Get()

		col := image.Uniform{color.RGBA{0, 30 + uint8(res.Resource*120), 0, 255}}
		draw.Draw(img, image.Rect(int(pos.X-rad), int(pos.Y-rad), int(pos.X+rad), int(pos.Y+rad)), &col, image.Point{}, draw.Src)
	}
}

// PostUpdateUI the system
func (s *DrawResources) PostUpdateUI(world *ecs.World) {}

// FinalizeUI the system
func (s *DrawResources) FinalizeUI(world *ecs.World) {}
