package main

import (
	"image"
	"image/color"
	"image/draw"

	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// DrawPatches system
type DrawPatches struct {
	canvas      generic.Resource[Image]
	patches     generic.Resource[Patches]
	patchFilter generic.Filter1[FlowerPatch]
}

// InitializeUI the system
func (s *DrawPatches) InitializeUI(world *ecs.World) {
	s.canvas = generic.NewResource[Image](world)
	s.patches = generic.NewResource[Patches](world)
	s.patchFilter = *generic.NewFilter1[FlowerPatch]()
}

// UpdateUI the system
func (s *DrawPatches) UpdateUI(world *ecs.World) {
	canvas := s.canvas.Get()
	img := canvas.Image
	cs := s.patches.Get().CellSize

	// Draw flower patches
	queryP := s.patchFilter.Query(world)
	for queryP.Next() {
		patch := queryP.Get()

		x := patch.X * cs
		y := patch.Y * cs
		col := image.Uniform{color.RGBA{0, 30 + uint8(patch.Resources*120), 0, 255}}
		draw.Draw(img, image.Rect(x, y, x+cs, y+cs), &col, image.Point{}, draw.Src)
	}
}

// PostUpdateUI the system
func (s *DrawPatches) PostUpdateUI(world *ecs.World) {}

// FinalizeUI the system
func (s *DrawPatches) FinalizeUI(world *ecs.World) {}
