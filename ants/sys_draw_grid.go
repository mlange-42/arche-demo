package ants

import (
	"image/color"

	"github.com/mlange-42/arche-demo/common"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// UISysDrawGrid is a system that draws grid nodes.
type UISysDrawGrid struct {
	canvas generic.Resource[common.Image]
	filter generic.Filter1[Position]
}

// InitializeUI the system
func (s *UISysDrawGrid) InitializeUI(world *ecs.World) {
	s.canvas = generic.NewResource[common.Image](world)
	s.filter = *generic.NewFilter1[Position]().With(generic.T[Node]())
}

// UpdateUI the system
func (s *UISysDrawGrid) UpdateUI(world *ecs.World) {
	canvas := s.canvas.Get()
	img := canvas.Image

	grey := color.RGBA{120, 120, 120, 255}

	query := s.filter.Query(world)
	for query.Next() {
		pos := query.Get()
		img.SetRGBA(int(pos.X), int(pos.Y), grey)
	}
}

// PostUpdateUI the system
func (s *UISysDrawGrid) PostUpdateUI(world *ecs.World) {}

// FinalizeUI the system
func (s *UISysDrawGrid) FinalizeUI(world *ecs.World) {}
