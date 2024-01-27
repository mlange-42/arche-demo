package boids

import (
	"math"

	"github.com/mlange-42/arche/ecs"
)

// Grid resource holding a grid of bins for lookup acceleration.
type Grid struct {
	Rows     int
	Cols     int
	CellSize int
	Cells    [][]ecs.Entity
}

// NewGrid creates a new Patches resource from world dimensions and grid cell size.
func NewGrid(width, height int, cellSize int) Grid {
	rows := height / cellSize
	cols := width / cellSize
	if height%cellSize != 0 {
		rows++
	}
	if width%cellSize != 0 {
		cols++
	}
	patches := make([][]ecs.Entity, cols)
	for i := 0; i < cols; i++ {
		patches[i] = make([]ecs.Entity, rows)
	}
	return Grid{
		Rows:     rows,
		Cols:     cols,
		CellSize: cellSize,
		Cells:    patches,
	}
}

// ToCell converts world coordinates to integer patch coordinates.
func (p *Grid) ToCell(x, y float64) (int, int) {
	cs := float64(p.CellSize)
	return int(x / cs), int(y / cs)
}

// ToCellCenter returns the world coordinates of the center of the cell
// the given point is in.
func (p *Grid) ToCellCenter(x, y float64) (float64, float64) {
	cs := float64(p.CellSize)
	return (math.Floor(x/cs) + 0.5) * cs,
		(math.Floor(y/cs) + 0.5) * cs
}
