package ants

import (
	"image"
	"image/color"
	"image/draw"

	"github.com/mlange-42/arche-demo/common"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// DrawGrid is a system that draws entities as white pixels on an [Image] resource.
type DrawGrid struct {
	canvas       generic.Resource[common.Image]
	nest         generic.Resource[Nest]
	nodeFilter   generic.Filter2[Position, NodeResource]
	scoutFilter  generic.Filter1[Position]
	forageFilter generic.Filter1[Position]
	returnFilter generic.Filter1[Position]

	posMap generic.Map1[Position]
}

// InitializeUI the system
func (s *DrawGrid) InitializeUI(world *ecs.World) {
	s.canvas = generic.NewResource[common.Image](world)
	s.nest = generic.NewResource[Nest](world)
	s.nodeFilter = *generic.NewFilter2[Position, NodeResource]()
	s.scoutFilter = *generic.NewFilter1[Position]().With(generic.T[ActScout]())
	s.forageFilter = *generic.NewFilter1[Position]().With(generic.T[ActForage]())
	s.returnFilter = *generic.NewFilter1[Position]().With(generic.T[ActReturn]())
}

// UpdateUI the system
func (s *DrawGrid) UpdateUI(world *ecs.World) {
	nest := s.nest.Get()

	black := image.Uniform{color.RGBA{0, 0, 0, 255}}
	grey := color.RGBA{100, 100, 100, 255}

	white := color.RGBA{255, 255, 255, 255}
	yellow := color.RGBA{255, 200, 0, 255}
	cyan := color.RGBA{0, 255, 255, 255}

	canvas := s.canvas.Get()
	img := canvas.Image
	draw.Draw(img, img.Bounds(), &black, image.Point{}, draw.Src)

	rad := 5.0

	draw.Draw(img, image.Rect(int(nest.Pos.X-rad), int(nest.Pos.Y-rad), int(nest.Pos.X+rad), int(nest.Pos.Y+rad)), &image.Uniform{grey}, image.Point{}, draw.Src)

	nodeQuery := s.nodeFilter.Query(world)
	for nodeQuery.Next() {
		pos, res := nodeQuery.Get()
		img.SetRGBA(int(pos.X), int(pos.Y), grey)

		col := image.Uniform{color.RGBA{0, 30 + uint8(res.Resource*120), 0, 255}}
		draw.Draw(img, image.Rect(int(pos.X-rad), int(pos.Y-rad), int(pos.X+rad), int(pos.Y+rad)), &col, image.Point{}, draw.Src)
	}

	forageQuery := s.forageFilter.Query(world)
	for forageQuery.Next() {
		pos := forageQuery.Get()
		img.SetRGBA(int(pos.X), int(pos.Y), white)
	}

	returnQuery := s.returnFilter.Query(world)
	for returnQuery.Next() {
		pos := returnQuery.Get()
		img.SetRGBA(int(pos.X), int(pos.Y), cyan)
	}

	scoutQuery := s.scoutFilter.Query(world)
	for scoutQuery.Next() {
		pos := scoutQuery.Get()
		img.SetRGBA(int(pos.X), int(pos.Y), yellow)
	}

	canvas.Redraw()
}

// PostUpdateUI the system
func (s *DrawGrid) PostUpdateUI(world *ecs.World) {}

// FinalizeUI the system
func (s *DrawGrid) FinalizeUI(world *ecs.World) {}
