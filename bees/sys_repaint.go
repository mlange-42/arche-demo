package bees

import (
	"image"
	"image/color"
	"image/draw"

	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// SysRepaint is a simple system that paints an [Image] resource to a [common.Canvas]
// and clears the image afterwards.
type SysRepaint struct {
	canvas generic.Resource[Image]
}

// InitializeUI the system
func (s *SysRepaint) InitializeUI(world *ecs.World) {
	s.canvas = generic.NewResource[Image](world)
}

// UpdateUI the system
func (s *SysRepaint) UpdateUI(world *ecs.World) {
	black := image.Uniform{color.RGBA{0, 0, 0, 255}}

	canvas := s.canvas.Get()
	img := canvas.Image

	canvas.Redraw()
	// Clear the image
	draw.Draw(img, img.Bounds(), &black, image.Point{}, draw.Src)

}

// PostUpdateUI the system
func (s *SysRepaint) PostUpdateUI(world *ecs.World) {}

// FinalizeUI the system
func (s *SysRepaint) FinalizeUI(world *ecs.World) {}
