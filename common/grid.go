package common

import "math"

// Grid data structure for 2D uniform grids.
type Grid[T any] struct {
	data     []T
	width    int
	height   int
	cellsize float64
}

// NewGrid returns a new Grid.
func NewGrid[T any](width, height int, cellsize ...float64) Grid[T] {
	cs := 1.0
	if len(cellsize) > 0 {
		cs = cellsize[0]
	}
	return Grid[T]{
		width:    width,
		height:   height,
		cellsize: cs,
		data:     make([]T, width*height),
	}
}

// NewGridExtent returns a new Grid with the given global extent.
func NewGridExtent[T any](width, height float64, cellsize ...float64) Grid[T] {
	cs := 1.0
	if len(cellsize) > 0 {
		cs = cellsize[0]
	}
	rows := int(height / cs)
	cols := int(width / cs)
	if math.Mod(height, cs) != 0 {
		rows++
	}
	if math.Mod(width, cs) != 0 {
		cols++
	}
	return NewGrid[T](cols, rows, cs)
}

// Get a value from the Grid.
func (g *Grid[T]) Get(x, y int) T {
	idx := x*g.height + y
	return g.data[idx]
}

// Get a pointer to a value from the Grid.
func (g *Grid[T]) GetPointer(x, y int) *T {
	idx := x*g.height + y
	return &g.data[idx]
}

// Set a value in the grid.
func (g *Grid[T]) Set(x, y int, value T) {
	idx := x*g.height + y
	g.data[idx] = value
}

// Fill the grid with a value.
func (g *Grid[T]) Fill(value T) {
	for i := range g.data {
		g.data[i] = value
	}
}

// Width of the Grid.
func (g *Grid[T]) Width() int {
	return g.width
}

// Height of the Grid.
func (g *Grid[T]) Height() int {
	return g.height
}

// Cell size of the Grid.
func (g *Grid[T]) Cellsize() float64 {
	return g.cellsize
}

// ToCell converts world coordinates to integer patch coordinates.
func (g *Grid[T]) ToCell(x, y float64) (int, int) {
	cs := float64(g.cellsize)
	return int(x / cs), int(y / cs)
}

// ToCellCenter returns the world coordinates of the center of the cell
// the given point is in.
func (g *Grid[T]) ToCellCenter(x, y float64) (float64, float64) {
	cs := float64(g.cellsize)
	return (math.Floor(x/cs) + 0.5) * cs,
		(math.Floor(y/cs) + 0.5) * cs
}

// CellCenter returns the world coordinates of the center of the given cell.
func (g *Grid[T]) CellCenter(x, y int) (float64, float64) {
	cs := float64(g.cellsize)
	return (float64(x) + 0.5) * cs,
		(float64(y) + 0.5) * cs
}

// Contains returns whether the grid contains the given cell.
func (g *Grid[T]) Contains(x, y int) bool {
	return x >= 0 && y >= 0 && x < g.width && y < g.height
}

// Contains returns whether the grid contains the given world coordinates.
func (g *Grid[T]) IsInBounds(x, y float64) bool {
	return x >= 0 && y >= 0 && x < float64(g.width)*g.cellsize && y < float64(g.height)*g.cellsize
}
