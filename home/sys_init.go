package main

import (
	"math/rand"

	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// InitEntities system
type InitEntities struct {
}

// Initialize the system
func (s *InitEntities) Initialize(world *ecs.World) {
	gridRes := generic.NewResource[Grid](world)
	grid := gridRes.Get()
	canvasRes := generic.NewResource[Image](world)
	canvas := canvasRes.Get()

	xOffset := (canvas.Width - grid.Width) / 2
	yOffset := (canvas.Height - grid.Height) / 2

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
			pos.X = rand.Float64() * float64(canvas.Width)
			pos.Y = rand.Float64() * float64(canvas.Height)
			targ.X = float64(x + xOffset)
			targ.Y = float64(y + yOffset)
		}
	}
	println(cnt, "entities")
}

// Update the system
func (s *InitEntities) Update(world *ecs.World) {}

// Finalize the system
func (s *InitEntities) Finalize(world *ecs.World) {}
