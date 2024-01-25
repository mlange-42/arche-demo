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
	canvas      generic.Resource[Image]
	patches     generic.Resource[Patches]
	hiveFilter  generic.Filter1[Position]
	patchFilter generic.Filter1[FlowerPatch]
	scoutFilter generic.Filter1[Position]
}

// InitializeUI the system
func (s *DrawHives) InitializeUI(world *ecs.World) {
	s.canvas = generic.NewResource[Image](world)
	s.patches = generic.NewResource[Patches](world)
	s.hiveFilter = *generic.NewFilter1[Position]().With(generic.T[Hive]())
	s.patchFilter = *generic.NewFilter1[FlowerPatch]()
	s.scoutFilter = *generic.NewFilter1[Position]().With(generic.T[HomeHive]())
}

// UpdateUI the system
func (s *DrawHives) UpdateUI(world *ecs.World) {
	black := image.Uniform{color.RGBA{0, 0, 0, 255}}
	blue := image.Uniform{color.RGBA{0, 0, 250, 255}}
	green := image.Uniform{color.RGBA{0, 120, 0, 255}}
	white := color.RGBA{255, 255, 255, 255}

	canvas := s.canvas.Get()
	img := canvas.Image
	cs := s.patches.Get().CellSize

	// Clear the image
	draw.Draw(img, img.Bounds(), &black, image.Point{}, draw.Src)

	// Draw flower patches
	queryP := s.patchFilter.Query(world)
	for queryP.Next() {
		pos := queryP.Get()

		x := pos.X * cs
		y := pos.Y * cs
		draw.Draw(img, image.Rect(x, y, x+cs, y+cs), &green, image.Point{}, draw.Src)
	}

	// Draw hives
	queryH := s.hiveFilter.Query(world)
	for queryH.Next() {
		pos := queryH.Get()

		draw.Draw(img, image.Rect(int(pos.X-2), int(pos.Y-2), int(pos.X+2), int(pos.Y+2)), &blue, image.Point{}, draw.Src)
	}

	// Draw scouts
	queryS := s.scoutFilter.Query(world)
	for queryS.Next() {
		pos := queryS.Get()

		img.SetRGBA(int(pos.X), int(pos.Y), white)
	}

	canvas.Redraw()
}

// PostUpdateUI the system
func (s *DrawHives) PostUpdateUI(world *ecs.World) {}

// FinalizeUI the system
func (s *DrawHives) FinalizeUI(world *ecs.World) {}
