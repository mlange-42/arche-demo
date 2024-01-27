package boids

import (
	"math/rand"

	"github.com/mlange-42/arche-demo/common"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// InitEntities is a system to initialize entities.
type InitEntities struct {
	Count int
}

// Initialize the system
func (s *InitEntities) Initialize(world *ecs.World) {
	canvasRes := generic.NewResource[common.Image](world)
	canvas := canvasRes.Get()

	builder := generic.NewMap2[Position, Velocity](world)

	query := builder.NewBatchQ(s.Count)
	for query.Next() {
		pos, vel := query.Get()
		pos.X = rand.Float64() * float64(canvas.Width)
		pos.Y = rand.Float64() * float64(canvas.Height)

		vel.X, vel.Y, _ = common.Norm(rand.NormFloat64(), rand.NormFloat64())
	}
}

// Update the system
func (s *InitEntities) Update(world *ecs.World) {}

// Finalize the system
func (s *InitEntities) Finalize(world *ecs.World) {}
