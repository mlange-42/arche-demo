package bees

import (
	"math/rand"

	"github.com/mlange-42/arche-demo/common"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// InitHives is a system that creates a number of randomly places hives.
type InitHives struct {
	// Target number of hives.
	Count int
}

// Initialize the system
func (s *InitHives) Initialize(world *ecs.World) {
	canvasRes := generic.NewResource[common.Image](world)
	canvas := canvasRes.Get()
	w := float64(canvas.Width)
	h := float64(canvas.Height)

	builder := generic.NewMap3[Position, Hive, Random256](world)

	query := builder.NewBatchQ(s.Count)
	for query.Next() {
		pos, _, r256 := query.Get()

		pos.X = rand.Float64()*w*0.8 + w*0.1
		pos.Y = rand.Float64()*h*0.8 + h*0.1
		r256.Value = uint8(rand.Int31n(256))
	}
}

// Update the system
func (s *InitHives) Update(world *ecs.World) {}

// Finalize the system
func (s *InitHives) Finalize(world *ecs.World) {}
