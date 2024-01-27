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
	grid  generic.Resource[Grid]
}

// Initialize the system
func (s *InitEntities) Initialize(world *ecs.World) {
	gridRes := generic.NewResource[Grid](world)
	grid := gridRes.Get()
	canvasRes := generic.NewResource[common.Image](world)
	canvas := canvasRes.Get()

	builder := generic.NewMap3[Position, Velocity, CurrentCell](world, generic.T[CurrentCell]())

	for i := 0; i < s.Count; i++ {
		x, y := rand.Float64()*float64(canvas.Width), rand.Float64()*float64(canvas.Height)
		vx, vy, _ := common.Norm(rand.NormFloat64(), rand.NormFloat64())

		cx, cy := grid.ToCell(x, y)
		cell := grid.Cells[cx][cy]

		_ = builder.NewWith(
			&Position{X: x, Y: y},
			&Velocity{X: vx, Y: vy},
			&CurrentCell{},
			cell,
		)
	}
}

// Update the system
func (s *InitEntities) Update(world *ecs.World) {}

// Finalize the system
func (s *InitEntities) Finalize(world *ecs.World) {}
