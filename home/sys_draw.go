package main

import (
	"image"
	"image/color"
	"image/draw"

	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// DrawEntities system
type DrawEntities struct {
	canvas generic.Resource[Image]
	filter generic.Filter1[Position]
}

// InitializeUI the system
func (s *DrawEntities) InitializeUI(world *ecs.World) {
	s.canvas = generic.NewResource[Image](world)
	s.filter = *generic.NewFilter1[Position]()
}

// UpdateUI the system
func (s *DrawEntities) UpdateUI(world *ecs.World) {
	black := color.RGBA{0, 0, 0, 255}
	white := color.RGBA{255, 255, 255, 255}

	canvas := s.canvas.Get()
	img := canvas.Image

	// Clear the image
	draw.Draw(img, img.Bounds(), &image.Uniform{black}, image.Point{}, draw.Src)

	// Draw pixel entities
	query := s.filter.Query(world)
	for query.Next() {
		pos := query.Get()

		img.SetRGBA(int(pos.X), int(pos.Y), white)
	}

	canvas.Redraw()
}

// PostUpdateUI the system
func (s *DrawEntities) PostUpdateUI(world *ecs.World) {}

// FinalizeUI the system
func (s *DrawEntities) FinalizeUI(world *ecs.World) {}
