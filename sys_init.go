package main

import (
	"math/rand"

	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// InitEntities system
type InitEntities struct {
	Count  int
	canvas generic.Resource[Canvas]
}

// Initialize the system
func (s *InitEntities) Initialize(world *ecs.World) {
	s.canvas = generic.NewResource[Canvas](world)

	mapper := generic.NewMap1[Position](world)
	query := mapper.NewBatchQ(s.Count)

	canvas := s.canvas.Get()
	w := canvas.Width
	h := canvas.Height
	for query.Next() {
		pos := query.Get()
		pos.X = rand.Float64() * w
		pos.Y = rand.Float64() * h
	}
}

// Update the system
func (s *InitEntities) Update(world *ecs.World) {}

// Finalize the system
func (s *InitEntities) Finalize(world *ecs.World) {}
