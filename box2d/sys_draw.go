package main

import (
	"image"
	"image/color"
	"image/draw"

	xdraw "golang.org/x/image/draw"

	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// DrawEntities system
type DrawEntities struct {
	canvas generic.Resource[Image]
	images generic.Resource[Images]
	filter generic.Filter1[Body]
}

// InitializeUI the system
func (s *DrawEntities) InitializeUI(world *ecs.World) {
	s.canvas = generic.NewResource[Image](world)
	s.images = generic.NewResource[Images](world)
	s.filter = *generic.NewFilter1[Body]()
}

// UpdateUI the system
func (s *DrawEntities) UpdateUI(world *ecs.World) {
	black := color.RGBA{0, 0, 0, 255}
	circle := s.images.Get().Circle
	bounds := circle.Bounds()

	canvas := s.canvas.Get()
	img := canvas.Image

	// Clear the image
	draw.Draw(img, img.Bounds(), &image.Uniform{black}, image.Point{}, draw.Src)

	// Draw pixel entities
	query := s.filter.Query(world)
	for query.Next() {
		bodyComp := query.Get()
		pos := bodyComp.Body.GetPosition()
		r := bodyComp.Radius
		rect := image.Rect(int(pos.X-r), int(pos.Y-r), int(pos.X+r), int(pos.Y+r))
		xdraw.ApproxBiLinear.Scale(img, rect, circle, bounds, draw.Over, nil)
	}

	canvas.Redraw()
}

// PostUpdateUI the system
func (s *DrawEntities) PostUpdateUI(world *ecs.World) {}

// FinalizeUI the system
func (s *DrawEntities) FinalizeUI(world *ecs.World) {}
