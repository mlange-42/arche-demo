package main

import (
	"image"
	"image/color"
	"image/draw"

	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/llgcode/draw2d/draw2dkit"
	"github.com/mlange-42/arche-demo/common"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// DrawEntities system
type DrawEntities struct {
	canvas generic.Resource[common.Canvas]
	filter generic.Filter1[Body]
}

// InitializeUI the system
func (s *DrawEntities) InitializeUI(world *ecs.World) {
	s.canvas = generic.NewResource[common.Canvas](world)
	s.filter = *generic.NewFilter1[Body]()
}

// UpdateUI the system
func (s *DrawEntities) UpdateUI(world *ecs.World) {
	black := color.RGBA{0, 0, 0, 255}
	white := color.RGBA{255, 255, 255, 255}

	canvas := s.canvas.Get()
	img := canvas.Image
	gc := draw2dimg.NewGraphicContext(img)

	// Clear the image
	draw.Draw(img, img.Bounds(), &image.Uniform{black}, image.Point{}, draw.Src)

	gc.SetStrokeColor(white)
	gc.SetLineWidth(1.2)

	// Draw pixel entities
	query := s.filter.Query(world)
	for query.Next() {
		bodyComp := query.Get()
		pos := bodyComp.Body.GetPosition()
		r := bodyComp.Radius

		draw2dkit.Circle(gc, pos.X, pos.Y, r)
		gc.Stroke()
	}

	canvas.Redraw()
}

// PostUpdateUI the system
func (s *DrawEntities) PostUpdateUI(world *ecs.World) {}

// FinalizeUI the system
func (s *DrawEntities) FinalizeUI(world *ecs.World) {}
