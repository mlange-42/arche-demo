package ants

import (
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// InitGrid is a system create a network on a grid.
type InitGrid struct{}

// Initialize the system
func (s *InitGrid) Initialize(world *ecs.World) {
	gridRes := generic.NewResource[Patches](world)
	grid := gridRes.Get()

	nodeBuilder := generic.NewMap2[Position, Node](world)
	//edgeBuilder := generic.NewMap1[Edge](world)

	query := nodeBuilder.NewBatchQ(grid.Cols * grid.Rows)
	cnt := 0
	for query.Next() {
		pos, _ := query.Get()
		i, j := cnt/grid.Rows, cnt%grid.Rows
		pos.X, pos.Y = grid.CellCenter(i, j)
		grid.Set(i, j, query.Entity())
		cnt++
	}

	_ = grid
}

// Update the system
func (s *InitGrid) Update(world *ecs.World) {}

// Finalize the system
func (s *InitGrid) Finalize(world *ecs.World) {}
