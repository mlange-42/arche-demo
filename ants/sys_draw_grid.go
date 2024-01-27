package ants

import (
	"image/color"

	"github.com/mlange-42/arche-demo/common"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// DrawGrid is a system that draws entities as white pixels on an [Image] resource.
type DrawGrid struct {
	canvas generic.Resource[common.Image]
	filter generic.Filter1[Position]
}

// InitializeUI the system
func (s *DrawGrid) InitializeUI(world *ecs.World) {
	s.canvas = generic.NewResource[common.Image](world)
	s.filter = *generic.NewFilter1[Position]().With(generic.T[Node]())
}

// UpdateUI the system
func (s *DrawGrid) UpdateUI(world *ecs.World) {
	white := color.RGBA{255, 255, 255, 255}

	canvas := s.canvas.Get()
	img := canvas.Image

	// Draw pixel entities
	query := s.filter.Query(world)
	for query.Next() {
		pos := query.Get()

		img.SetRGBA(int(pos.X), int(pos.Y), white)
	}
}

// PostUpdateUI the system
func (s *DrawGrid) PostUpdateUI(world *ecs.World) {}

// FinalizeUI the system
func (s *DrawGrid) FinalizeUI(world *ecs.World) {}
