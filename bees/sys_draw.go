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

	followFilter generic.Filter1[Position]
	scoutFilter  generic.Filter1[Position]
	forageFilter generic.Filter1[Position]
	returnFilter generic.Filter1[Position]
	inHiveFilter generic.Filter1[Position]
	waggleFilter generic.Filter1[Position]
}

// InitializeUI the system
func (s *DrawHives) InitializeUI(world *ecs.World) {
	s.canvas = generic.NewResource[Image](world)
	s.patches = generic.NewResource[Patches](world)
	s.hiveFilter = *generic.NewFilter1[Position]().With(generic.T[Hive]())
	s.patchFilter = *generic.NewFilter1[FlowerPatch]()

	s.followFilter = *generic.NewFilter1[Position]().With(generic.T[ActFollow]())
	s.scoutFilter = *generic.NewFilter1[Position]().With(generic.T[ActScout]())
	s.forageFilter = *generic.NewFilter1[Position]().With(generic.T[ActForage]())
	s.returnFilter = *generic.NewFilter1[Position]().With(generic.T[ActReturn]())
	s.inHiveFilter = *generic.NewFilter1[Position]().With(generic.T[ActInHive]())
	s.waggleFilter = *generic.NewFilter1[Position]().With(generic.T[ActWaggleDance]())
}

// UpdateUI the system
func (s *DrawHives) UpdateUI(world *ecs.World) {
	black := image.Uniform{color.RGBA{0, 0, 0, 255}}
	blue := image.Uniform{color.RGBA{0, 0, 250, 255}}

	followCol := color.RGBA{255, 255, 255, 255}
	scoutCol := color.RGBA{255, 255, 100, 255}
	forageCol := color.RGBA{255, 0, 255, 255}
	returnCol := color.RGBA{0, 255, 255, 255}
	inHiveCol := color.RGBA{100, 100, 255, 255}
	waggleCol := color.RGBA{255, 50, 50, 255}

	canvas := s.canvas.Get()
	img := canvas.Image
	cs := s.patches.Get().CellSize

	// Clear the image
	draw.Draw(img, img.Bounds(), &black, image.Point{}, draw.Src)

	// Draw flower patches
	queryP := s.patchFilter.Query(world)
	for queryP.Next() {
		patch := queryP.Get()

		x := patch.X * cs
		y := patch.Y * cs
		col := image.Uniform{color.RGBA{0, 30 + uint8(patch.Resources*120), 0, 255}}
		draw.Draw(img, image.Rect(x, y, x+cs, y+cs), &col, image.Point{}, draw.Src)
	}

	// Draw hives
	queryH := s.hiveFilter.Query(world)
	for queryH.Next() {
		pos := queryH.Get()

		draw.Draw(img, image.Rect(int(pos.X-2), int(pos.Y-2), int(pos.X+2), int(pos.Y+2)), &blue, image.Point{}, draw.Src)
	}

	queryFollow := s.followFilter.Query(world)
	for queryFollow.Next() {
		pos := queryFollow.Get()
		img.SetRGBA(int(pos.X), int(pos.Y), followCol)
	}
	queryScouts := s.scoutFilter.Query(world)
	for queryScouts.Next() {
		pos := queryScouts.Get()
		img.SetRGBA(int(pos.X), int(pos.Y), scoutCol)
	}
	queryForage := s.forageFilter.Query(world)
	for queryForage.Next() {
		pos := queryForage.Get()
		img.SetRGBA(int(pos.X), int(pos.Y), forageCol)
	}
	queryReturn := s.returnFilter.Query(world)
	for queryReturn.Next() {
		pos := queryReturn.Get()
		img.SetRGBA(int(pos.X), int(pos.Y), returnCol)
	}
	queryInHive := s.inHiveFilter.Query(world)
	for queryInHive.Next() {
		pos := queryInHive.Get()
		img.SetRGBA(int(pos.X), int(pos.Y), inHiveCol)
	}
	queryWaggle := s.waggleFilter.Query(world)
	for queryWaggle.Next() {
		pos := queryWaggle.Get()
		img.SetRGBA(int(pos.X), int(pos.Y), waggleCol)
	}

	canvas.Redraw()
}

// PostUpdateUI the system
func (s *DrawHives) PostUpdateUI(world *ecs.World) {}

// FinalizeUI the system
func (s *DrawHives) FinalizeUI(world *ecs.World) {}
