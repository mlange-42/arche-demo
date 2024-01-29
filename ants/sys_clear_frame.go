package ants

import (
	"image"
	"image/color"
	"image/draw"

	"github.com/mlange-42/arche-demo/common"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// SysClearFrame is a simple system that clears the [Image] resource
// before other systems draw on it.
type SysClearFrame struct {
	canvas generic.Resource[common.Image]
}

// InitializeUI the system
func (s *SysClearFrame) InitializeUI(world *ecs.World) {
	s.canvas = generic.NewResource[common.Image](world)
}

// UpdateUI the system
func (s *SysClearFrame) UpdateUI(world *ecs.World) {
	black := image.Uniform{color.RGBA{0, 0, 0, 255}}

	canvas := s.canvas.Get()
	img := canvas.Image

	// Clear the image
	draw.Draw(img, img.Bounds(), &black, image.Point{}, draw.Src)
}

// PostUpdateUI the system
func (s *SysClearFrame) PostUpdateUI(world *ecs.World) {}

// FinalizeUI the system
func (s *SysClearFrame) FinalizeUI(world *ecs.World) {}
