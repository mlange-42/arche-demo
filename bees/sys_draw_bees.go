package bees

import (
	"github.com/mlange-42/arche-demo/common"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// DrawBees is a system for drawing bees as colored pixels.
type DrawBees struct {
	canvas generic.Resource[common.Image]
	colors generic.Resource[Colors]

	followFilter generic.Filter1[Position]
	scoutFilter  generic.Filter1[Position]
	forageFilter generic.Filter1[Position]
	returnFilter generic.Filter1[Position]
	inHiveFilter generic.Filter1[Position]
	waggleFilter generic.Filter1[Position]
}

// InitializeUI the system
func (s *DrawBees) InitializeUI(world *ecs.World) {
	s.canvas = generic.NewResource[common.Image](world)
	s.colors = generic.NewResource[Colors](world)

	s.followFilter = *generic.NewFilter1[Position]().With(generic.T[ActFollow]())
	s.scoutFilter = *generic.NewFilter1[Position]().With(generic.T[ActScout]())
	s.forageFilter = *generic.NewFilter1[Position]().With(generic.T[ActForage]())
	s.returnFilter = *generic.NewFilter1[Position]().With(generic.T[ActReturn]())
	s.inHiveFilter = *generic.NewFilter1[Position]().With(generic.T[ActInHive]())
	s.waggleFilter = *generic.NewFilter1[Position]().With(generic.T[ActWaggleDance]())
}

// UpdateUI the system
func (s *DrawBees) UpdateUI(world *ecs.World) {
	cols := s.colors.Get()

	canvas := s.canvas.Get()
	img := canvas.Image

	queryFollow := s.followFilter.Query(world)
	for queryFollow.Next() {
		pos := queryFollow.Get()
		img.SetRGBA(int(pos.X), int(pos.Y), cols.Follow)
	}
	queryScouts := s.scoutFilter.Query(world)
	for queryScouts.Next() {
		pos := queryScouts.Get()
		img.SetRGBA(int(pos.X), int(pos.Y), cols.Scout)
	}
	queryForage := s.forageFilter.Query(world)
	for queryForage.Next() {
		pos := queryForage.Get()
		img.SetRGBA(int(pos.X), int(pos.Y), cols.Forage)
	}
	queryReturn := s.returnFilter.Query(world)
	for queryReturn.Next() {
		pos := queryReturn.Get()
		img.SetRGBA(int(pos.X), int(pos.Y), cols.Return)
	}
	/*
		queryInHive := s.inHiveFilter.Query(world)
		for queryInHive.Next() {
			pos := queryInHive.Get()
			img.SetRGBA(int(pos.X), int(pos.Y), cols.InHive)
		}
		queryWaggle := s.waggleFilter.Query(world)
		for queryWaggle.Next() {
			pos := queryWaggle.Get()
			img.SetRGBA(int(pos.X), int(pos.Y), cols.Waggle)
		}
	*/
}

// PostUpdateUI the system
func (s *DrawBees) PostUpdateUI(world *ecs.World) {}

// FinalizeUI the system
func (s *DrawBees) FinalizeUI(world *ecs.World) {}
