package ants

import (
	"image/color"

	"github.com/mlange-42/arche-demo/common"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// DrawAnts is a system that draws ants.
type DrawAnts struct {
	canvas       generic.Resource[common.Image]
	scoutFilter  generic.Filter1[Position]
	forageFilter generic.Filter1[Position]
	returnFilter generic.Filter1[Position]
}

// InitializeUI the system
func (s *DrawAnts) InitializeUI(world *ecs.World) {
	s.canvas = generic.NewResource[common.Image](world)
	s.scoutFilter = *generic.NewFilter1[Position]().With(generic.T[ActScout]())
	s.forageFilter = *generic.NewFilter1[Position]().With(generic.T[ActForage]())
	s.returnFilter = *generic.NewFilter1[Position]().With(generic.T[ActReturn]())
}

// UpdateUI the system
func (s *DrawAnts) UpdateUI(world *ecs.World) {
	canvas := s.canvas.Get()
	img := canvas.Image

	white := color.RGBA{255, 255, 255, 255}
	yellow := color.RGBA{255, 200, 0, 255}
	cyan := color.RGBA{0, 255, 255, 255}

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
}

// PostUpdateUI the system
func (s *DrawAnts) PostUpdateUI(world *ecs.World) {}

// FinalizeUI the system
func (s *DrawAnts) FinalizeUI(world *ecs.World) {}
