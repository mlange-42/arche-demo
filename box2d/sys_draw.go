package main

import (
	"image"
	"image/color"
	"image/draw"
	"math"

	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// DrawEntities system
type DrawEntities struct {
	canvas generic.Resource[Image]
	filter generic.Filter1[Body]
}

// InitializeUI the system
func (s *DrawEntities) InitializeUI(world *ecs.World) {
	s.canvas = generic.NewResource[Image](world)
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
	gc.SetLineWidth(1)

	// Draw pixel entities
	query := s.filter.Query(world)
	for query.Next() {
		bodyComp := query.Get()
		pos := bodyComp.Body.GetPosition()
		ang := bodyComp.Body.GetAngle()
		r := bodyComp.Radius

		gc.BeginPath()
		gc.MoveTo(pos.X, pos.Y)
		gc.ArcTo(pos.X, pos.Y, r, r, ang, -math.Pi*2)
		gc.Close()
		gc.Stroke()
	}

	canvas.Redraw()
}

// PostUpdateUI the system
func (s *DrawEntities) PostUpdateUI(world *ecs.World) {}

// FinalizeUI the system
func (s *DrawEntities) FinalizeUI(world *ecs.World) {}
