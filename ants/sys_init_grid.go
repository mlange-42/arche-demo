package ants

import (
	"github.com/mlange-42/arche-demo/common"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// coord helper struct
type coord struct {
	X int
	Y int
}

// InitGrid is a system create a network on a grid.
type InitGrid struct{}

// Initialize the system
func (s *InitGrid) Initialize(world *ecs.World) {
	gridRes := generic.NewResource[Patches](world)
	grid := gridRes.Get()

	nodeBuilder := generic.NewMap2[Position, Node](world)
	edgeBuilder := generic.NewMap1[Edge](world)

	query := nodeBuilder.NewBatchQ(grid.Cols * grid.Rows)
	cnt := 0
	for query.Next() {
		pos, _ := query.Get()
		i, j := cnt/grid.Rows, cnt%grid.Rows
		pos.X, pos.Y = grid.CellCenter(i, j)
		grid.Set(i, j, query.Entity())
		cnt++
	}

	dirs := []coord{
		{1, 0},
		{1, 1},
		{0, 1},
		{-1, 1},
	}

	for x := 0; x < grid.Cols-1; x++ {
		for y := 0; y < grid.Cols-1; y++ {
			thisEntity := grid.Get(x, y)
			thisPos, thisNode := nodeBuilder.Get(thisEntity)
			for _, off := range dirs {
				xx, yy := x+off.X, y+off.Y
				if !grid.Contains(xx, yy) {
					continue
				}
				otherEntity := grid.Get(xx, yy)
				otherPos, otherNode := nodeBuilder.Get(otherEntity)

				ln := common.Distance(thisPos.X, thisPos.Y, otherPos.X, otherPos.Y)
				edge1 := edgeBuilder.NewWith(&Edge{From: thisEntity, To: otherEntity, Length: ln})
				edge2 := edgeBuilder.NewWith(&Edge{From: otherEntity, To: thisEntity, Length: ln})

				thisNode.Add(edge2, edge1)
				otherNode.Add(edge1, edge2)
			}
		}
	}

	_ = grid
}

// Update the system
func (s *InitGrid) Update(world *ecs.World) {}

// Finalize the system
func (s *InitGrid) Finalize(world *ecs.World) {}
