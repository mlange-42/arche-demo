package main

import (
	"math/rand"

	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// InitEntities system
type InitEntities struct {
	grid generic.Resource[Grid]
}

// Initialize the system
func (s *InitEntities) Initialize(world *ecs.World) {
	s.grid = generic.NewResource[Grid](world)

	grid := s.grid.Get()
	builder := generic.NewMap3[Position, Velocity, Target](world)

	cnt := 0
	for y := 0; y < grid.Height; y++ {
		for x := 0; x < grid.Width; x++ {
			if !grid.Data[y][x] {
				continue
			}
			cnt++
			e := builder.New()
			pos, _, targ := builder.Get(e)
			pos.X = rand.Float64() * float64(grid.Width)
			pos.Y = rand.Float64() * float64(grid.Height)
			targ.X = float64(x)
			targ.Y = float64(y)
		}
	}
	println(cnt, "entities")
}

// Update the system
func (s *InitEntities) Update(world *ecs.World) {}

// Finalize the system
func (s *InitEntities) Finalize(world *ecs.World) {}
