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

// InitGrid is a system to create a network on a grid.
type InitGrid struct{}

// Initialize the system
func (s *InitGrid) Initialize(world *ecs.World) {
	gridRes := generic.NewResource[Patches](world)
	grid := gridRes.Get()

	nodeBuilder := generic.NewMap2[Position, Node](world)
	edgeBuilder := generic.NewMap3[Edge, EdgeGeometry, Trace](world)

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

				dx, dy, ln := common.Norm(otherPos.X-thisPos.X, otherPos.Y-thisPos.Y)

				edge1 := edgeBuilder.NewWith(
					&Edge{From: thisEntity, To: otherEntity},
					&EdgeGeometry{From: *thisPos, Dir: Position{X: dx, Y: dy}, Length: ln},
					&Trace{},
				)
				edge2 := edgeBuilder.NewWith(
					&Edge{From: otherEntity, To: thisEntity},
					&EdgeGeometry{From: *otherPos, Dir: Position{X: -dx, Y: -dy}, Length: ln},
					&Trace{},
				)

				thisNode.Add(edge2, edge1)
				otherNode.Add(edge1, edge2)
			}
		}
	}
}

// Update the system
func (s *InitGrid) Update(world *ecs.World) {}

// Finalize the system
func (s *InitGrid) Finalize(world *ecs.World) {}
