package ants

import (
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// SysInitNest is a system to initialize an ant nest, and to populate it with ants.
type SysInitNest struct {
	AntsPerNest int
	nest        Nest
}

// Initialize the system
func (s *SysInitNest) Initialize(world *ecs.World) {
	gridRes := generic.NewResource[Patches](world)
	grid := gridRes.Get()

	node := grid.Get(grid.Cols/2, grid.Rows/2)
	posMap := generic.NewMap1[Position](world)
	nestPos := posMap.Get(node)

	s.nest = Nest{
		Node: node,
		Pos:  *nestPos,
	}
	ecs.AddResource(world, &s.nest)

	antBuilder := generic.NewMap3[Position, ActInNest, Ant](world)

	query := antBuilder.NewBatchQ(s.AntsPerNest)
	for query.Next() {
		pos, _, _ := query.Get()
		pos.X, pos.Y = nestPos.X, nestPos.Y
	}
}

// Update the system
func (s *SysInitNest) Update(world *ecs.World) {}

// Finalize the system
func (s *SysInitNest) Finalize(world *ecs.World) {}
