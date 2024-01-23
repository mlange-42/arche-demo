package main

import (
	"image"
	"image/color"
	"image/draw"

	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// DrawEntities system
type DrawEntities struct {
	canvas generic.Resource[Canvas]
	filter generic.Filter1[Position]
	image  *image.RGBA
}

// InitializeUI the system
func (s *DrawEntities) InitializeUI(world *ecs.World) {
	s.canvas = generic.NewResource[Canvas](world)
	s.filter = *generic.NewFilter1[Position]()

	c := s.canvas.Get()
	s.image = image.NewRGBA(image.Rect(0, 0, int(c.Width), int(c.Height)))
}

// UpdateUI the system
func (s *DrawEntities) UpdateUI(world *ecs.World) {

	gc := s.canvas.Get().Canvas.Gc()
	gc.Filter = draw2dimg.LinearFilter

	black := color.RGBA{0, 0, 0, 255}
	white := color.RGBA{255, 255, 255, 255}
	draw.Draw(s.image, s.image.Bounds(), &image.Uniform{black}, image.Point{}, draw.Src)

	//gc.SetFillColor(color.RGBA{0x00, 0x00, 0x00, 0xff})
	//gc.Clear()

	//gc.SetFillColor(color.RGBA{0xff, 0xff, 0xff, 0xff})
	//gc.SetStrokeColor(color.RGBA{0xff, 0xff, 0xff, 0xff})

	query := s.filter.Query(world)
	for query.Next() {
		pos := query.Get()

		s.image.SetRGBA(int(pos.X), int(pos.Y), white)
	}

	gc.DrawImage(s.image)
}

// PostUpdateUI the system
func (s *DrawEntities) PostUpdateUI(world *ecs.World) {}

// FinalizeUI the system
func (s *DrawEntities) FinalizeUI(world *ecs.World) {}
