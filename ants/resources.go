package ants

import (
	"math"

	"github.com/mlange-42/arche/ecs"
)

// Patches resource holding a grid of potential flower patches.
type Patches struct {
	Rows     int
	Cols     int
	CellSize int
	nodes    []ecs.Entity
}

// NewPatches creates a new Patches resource from world dimensions and grid cell size.
func NewPatches(width, height int, cellSize int) Patches {
	rows := height / cellSize
	cols := width / cellSize
	if height%cellSize != 0 {
		rows++
	}
	if width%cellSize != 0 {
		cols++
	}
	return Patches{
		Rows:     rows,
		Cols:     cols,
		CellSize: cellSize,
		nodes:    make([]ecs.Entity, cols*rows),
	}
}

// Get the node at a coordinate.
func (p *Patches) Get(x, y int) ecs.Entity {
	return p.nodes[x*p.Rows+y]
}

// Set the node at a coordinate.
func (p *Patches) Set(x, y int, node ecs.Entity) {
	p.nodes[x*p.Rows+y] = node
}

// ToCell converts world coordinates to integer patch coordinates.
func (p *Patches) ToCell(x, y float64) (int, int) {
	cs := float64(p.CellSize)
	return int(x / cs), int(y / cs)
}

// ToCellCenter returns the world coordinates of the center of the cell
// the given point is in.
func (p *Patches) ToCellCenter(x, y float64) (float64, float64) {
	cs := float64(p.CellSize)
	return (math.Floor(x/cs) + 0.5) * cs,
		(math.Floor(y/cs) + 0.5) * cs
}

// CellCenter returns the world coordinates of the center of the given cell.
func (p *Patches) CellCenter(x, y int) (float64, float64) {
	cs := float64(p.CellSize)
	return (float64(x) + 0.5) * cs,
		(float64(y) + 0.5) * cs
}

// Contains returns whether the grid contains the given cell.
func (p *Patches) Contains(x, y int) bool {
	return x >= 0 && y >= 0 && x < p.Cols && y < p.Rows
}
