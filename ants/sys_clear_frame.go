package ants

import (
	"image"
	"image/color"
	"image/draw"

	"github.com/mlange-42/arche-demo/common"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// UISysClearFrame is a simple system that clears the [Image] resource
// before other systems draw on it.
type UISysClearFrame struct {
	canvas generic.Resource[common.Image]
}

// InitializeUI the system
func (s *UISysClearFrame) InitializeUI(world *ecs.World) {
	s.canvas = generic.NewResource[common.Image](world)
}

// UpdateUI the system
func (s *UISysClearFrame) UpdateUI(world *ecs.World) {
	black := image.Uniform{color.RGBA{0, 0, 0, 255}}

	canvas := s.canvas.Get()
	img := canvas.Image

	// Clear the image
	draw.Draw(img, img.Bounds(), &black, image.Point{}, draw.Src)
}

// PostUpdateUI the system
func (s *UISysClearFrame) PostUpdateUI(world *ecs.World) {}

// FinalizeUI the system
func (s *UISysClearFrame) FinalizeUI(world *ecs.World) {}
