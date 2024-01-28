package ants

import (
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// InitNest is a system to initialize an ant nest, and to populate it with ants.
type InitNest struct {
	AntsPerNest int
	nest        Nest
}

// Initialize the system
func (s *InitNest) Initialize(world *ecs.World) {
	gridRes := generic.NewResource[Patches](world)
	grid := gridRes.Get()

	node := grid.Get(grid.Cols/2, grid.Rows/2)
	s.nest = Nest{
		Node: node,
	}
	ecs.AddResource(world, &s.nest)

	posMap := generic.NewMap1[Position](world)
	nestPos := posMap.Get(node)

	antBuilder := generic.NewMap3[Position, ActInNest, Ant](world)

	query := antBuilder.NewBatchQ(s.AntsPerNest)
	for query.Next() {
		pos, _, _ := query.Get()
		pos.X, pos.Y = nestPos.X, nestPos.Y
	}
}

// Update the system
func (s *InitNest) Update(world *ecs.World) {}

// Finalize the system
func (s *InitNest) Finalize(world *ecs.World) {}
