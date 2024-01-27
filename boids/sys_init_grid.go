package boids

import (
	"github.com/mlange-42/arche-demo/common"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// InitGrid is a system to initialize the acceleration grid.
type InitGrid struct {
	CellSize int
	grid     Grid
}

// Initialize the system
func (s *InitGrid) Initialize(world *ecs.World) {
	canvasRes := generic.NewResource[common.Image](world)
	canvas := canvasRes.Get()

	s.grid = NewGrid(canvas.Width, canvas.Height, s.CellSize)
	ecs.AddResource(world, &s.grid)

	builder := generic.NewMap1[Cell](world)

	for i := 0; i < s.grid.Cols; i++ {
		for j := 0; j < s.grid.Rows; j++ {
			s.grid.Cells[i][j] = builder.NewWith(&Cell{X: i, Y: j})
		}
	}
}

// Update the system
func (s *InitGrid) Update(world *ecs.World) {}

// Finalize the system
func (s *InitGrid) Finalize(world *ecs.World) {}
