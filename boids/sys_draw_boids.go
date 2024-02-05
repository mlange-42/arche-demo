package boids

import (
	"image"
	"image/color"
	"image/draw"

	"github.com/mlange-42/arche-demo/common"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// UISysDrawBoids is a system that draws ants.
type UISysDrawBoids struct {
	canvas generic.Resource[common.EbitenImage]
	filter generic.Filter1[Position]

	image *image.RGBA
}

// InitializeUI the system
func (s *UISysDrawBoids) InitializeUI(world *ecs.World) {
	s.canvas = generic.NewResource[common.EbitenImage](world)
	s.filter = *generic.NewFilter1[Position]()

	screen := s.canvas.Get()
	s.image = image.NewRGBA(screen.Image.Bounds())
}

// UpdateUI the system
func (s *UISysDrawBoids) UpdateUI(world *ecs.World) {
	canvas := s.canvas.Get()
	img := canvas.Image

	black := color.RGBA{0, 0, 0, 255}
	white := color.RGBA{255, 255, 255, 255}

	// Clear the image
	draw.Draw(s.image, s.image.Bounds(), &image.Uniform{black}, image.Point{}, draw.Src)

	query := s.filter.Query(world)
	for query.Next() {
		pos := query.Get()
		s.image.SetRGBA(int(pos.X), int(pos.Y), white)
	}

	img.WritePixels(s.image.Pix)
}

// PostUpdateUI the system
func (s *UISysDrawBoids) PostUpdateUI(world *ecs.World) {}

// FinalizeUI the system
func (s *UISysDrawBoids) FinalizeUI(world *ecs.World) {}
