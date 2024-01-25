package main

import (
	"math/rand"

	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// InitHives system
type InitHives struct {
	Count int
}

// Initialize the system
func (s *InitHives) Initialize(world *ecs.World) {
	canvasRes := generic.NewResource[Image](world)
	canvas := canvasRes.Get()
	w := float64(canvas.Width)
	h := float64(canvas.Height)

	builder := generic.NewMap2[Position, Hive](world)

	query := builder.NewBatchQ(s.Count)
	for query.Next() {
		pos, _ := query.Get()

		pos.X = rand.Float64()*w*0.8 + w*0.1
		pos.Y = rand.Float64()*h*0.8 + h*0.1
	}
}

// Update the system
func (s *InitHives) Update(world *ecs.World) {}

// Finalize the system
func (s *InitHives) Finalize(world *ecs.World) {}
