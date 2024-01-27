package main

import (
	"image"
	"image/color"
	"image/draw"

	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// DrawHives system
type DrawHives struct {
	canvas     generic.Resource[Image]
	hiveFilter generic.Filter1[Position]

	scoutFilter  generic.Filter0
	followFilter generic.Filter0
	forageFilter generic.Filter0
	returnFilter generic.Filter0
	waggleFilter generic.Filter0
	idleFilter   generic.Filter0
}

// InitializeUI the system
func (s *DrawHives) InitializeUI(world *ecs.World) {
	s.canvas = generic.NewResource[Image](world)
	s.hiveFilter = *generic.NewFilter1[Position]().With(generic.T[Hive]())
}

// UpdateUI the system
func (s *DrawHives) UpdateUI(world *ecs.World) {
	blue := image.Uniform{color.RGBA{0, 0, 250, 255}}

	canvas := s.canvas.Get()
	img := canvas.Image

	// Draw hives
	queryH := s.hiveFilter.Query(world)
	for queryH.Next() {
		pos := queryH.Get()

		draw.Draw(img, image.Rect(int(pos.X-2), int(pos.Y-2), int(pos.X+2), int(pos.Y+2)), &blue, image.Point{}, draw.Src)
	}
}

// PostUpdateUI the system
func (s *DrawHives) PostUpdateUI(world *ecs.World) {}

// FinalizeUI the system
func (s *DrawHives) FinalizeUI(world *ecs.World) {}