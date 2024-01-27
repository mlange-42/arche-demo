package boids

import (
	"fmt"
	"math"
)

// Grid resource holding a grid of bins for lookup acceleration.
type Grid struct {
	Rows      int
	Cols      int
	CellSize  int
	maxPerBin uint32
	perCell   []uint32
	cells     []GridEntry
}

// GridEntry for the acceleration grid.
type GridEntry struct {
	X  float64
	Y  float64
	VX float64
	VY float64
}

// Set the values of a grid entry
func (e *GridEntry) Set(pos *Position, vel *Velocity) {
	e.X = pos.X
	e.Y = pos.Y
	e.VX = vel.X
	e.VY = vel.Y
}

// NewGrid creates a new Patches resource from world dimensions and grid cell size.
func NewGrid(width, height int, cellSize int, maxPerBin int) Grid {
	rows := height / cellSize
	cols := width / cellSize
	if height%cellSize != 0 {
		rows++
	}
	if width%cellSize != 0 {
		cols++
	}

	return Grid{
		Rows:      rows,
		Cols:      cols,
		CellSize:  cellSize,
		maxPerBin: uint32(maxPerBin),
		perCell:   make([]uint32, rows*cols),
		cells:     make([]GridEntry, rows*cols*maxPerBin),
	}
}

// Clear all cell.
func (g *Grid) Clear() {
	for i := 0; i < len(g.perCell); i++ {
		g.perCell[i] = 0
	}
}

// Add an entry to a cell.
func (g *Grid) Add(x, y int, pos *Position, vel *Velocity) {
	idx := uint32(x*g.Rows + y)
	count := g.perCell[idx]
	if count+1 >= g.maxPerBin {
		panic(fmt.Sprintf("grid cell is full: %d %d", x, y))
	}
	g.cells[idx*g.maxPerBin+count].Set(pos, vel)
}

// Get all entries of a cell.
func (g *Grid) Get(x, y int) []GridEntry {
	idx := uint32(x*g.Rows + y)
	count := g.perCell[idx]
	return g.cells[idx*g.maxPerBin : idx*g.maxPerBin+count]
}

// ToCell converts world coordinates to integer patch coordinates.
func (g *Grid) ToCell(x, y float64) (int, int) {
	cs := float64(g.CellSize)
	return int(x / cs), int(y / cs)
}

// ToCellCenter returns the world coordinates of the center of the cell
// the given point is in.
func (g *Grid) ToCellCenter(x, y float64) (float64, float64) {
	cs := float64(g.CellSize)
	return (math.Floor(x/cs) + 0.5) * cs,
		(math.Floor(y/cs) + 0.5) * cs
}
