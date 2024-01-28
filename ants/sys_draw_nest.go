package ants

import (
	"image"
	"image/color"
	"image/draw"

	"github.com/mlange-42/arche-demo/common"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// DrawNest is a system that draws the ant nest.
type DrawNest struct {
	canvas generic.Resource[common.Image]
	nest   generic.Resource[Nest]
}

// InitializeUI the system
func (s *DrawNest) InitializeUI(world *ecs.World) {
	s.canvas = generic.NewResource[common.Image](world)
	s.nest = generic.NewResource[Nest](world)
}

// UpdateUI the system
func (s *DrawNest) UpdateUI(world *ecs.World) {
	nest := s.nest.Get()
	canvas := s.canvas.Get()
	img := canvas.Image

	grey := color.RGBA{100, 100, 100, 255}
	rad := 5.0

	draw.Draw(
		img,
		image.Rect(int(nest.Pos.X-rad), int(nest.Pos.Y-rad), int(nest.Pos.X+rad), int(nest.Pos.Y+rad)),
		&image.Uniform{grey}, image.Point{}, draw.Src,
	)
}

// PostUpdateUI the system
func (s *DrawNest) PostUpdateUI(world *ecs.World) {}

// FinalizeUI the system
func (s *DrawNest) FinalizeUI(world *ecs.World) {}
