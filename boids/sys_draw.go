package boids

import (
	"image"
	"image/color"
	"image/draw"

	"github.com/mlange-42/arche-demo/common"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// DrawEntities is a system that draws entities as white pixels on an [Image] resource.
type DrawEntities struct {
	canvas generic.Resource[common.Image]
	grid   generic.Resource[Grid]
	filter generic.Filter1[Position]
}

// InitializeUI the system
func (s *DrawEntities) InitializeUI(world *ecs.World) {
	s.canvas = generic.NewResource[common.Image](world)
	s.grid = generic.NewResource[Grid](world)
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

	/*
		grid := s.grid.Get()
		cs := grid.CellSize
		for i := 0; i < grid.Cols; i++ {
			for j := 0; j < grid.Rows; j++ {
				cnt := grid.Count(i, j)
				draw.Draw(img, image.Rect(i*cs, j*cs, i*cs+cs, j*cs+cs), &image.Uniform{color.RGBA{0, uint8(cnt * 8), 0, 255}}, image.Point{}, draw.Src)
			}
		}
	*/

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
