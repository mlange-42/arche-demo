package main

import (
	"image"

	"github.com/mlange-42/arche/ecs"
)

// Position component.
type Position struct {
	X float64
	Y float64
}

// Direction component.
type Direction struct {
	X float64
	Y float64
}

// HomeHive component.
type HomeHive struct {
	ecs.Relation
}

// ActScout component.
type ActScout struct {
	Start int64
}

// ActForage activity component.
type ActForage struct {
	Start int64
}

// Random256 contains an uint8 value for scheduling things in intervals, but randomized over entities.
type Random256 struct {
	Value uint8
}

// Hive component.
type Hive struct{}

// FlowerPatch component
type FlowerPatch struct {
	X int
	Y int
}

// Image resource.
type Image struct {
	Image  *image.RGBA
	Width  int
	Height int
	Redraw func()
}

// Params resource
type Params struct {
	MaxBeeSpeed float64
}

// Patches resource holding the dimensions of the landscape grid.
type Patches struct {
	Rows     int
	Cols     int
	CellSize int
	Patches  [][]ecs.Entity
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
	patches := make([][]ecs.Entity, cols)
	for i := 0; i < cols; i++ {
		patches[i] = make([]ecs.Entity, rows)
	}
	return Patches{
		Rows:     rows,
		Cols:     cols,
		CellSize: cellSize,
		Patches:  patches,
	}
}

// ToCell converts world coordinates to integer patch coordinates.
func (p *Patches) ToCell(x, y float64) (int, int) {
	return int(x / float64(p.CellSize)), int(y / float64(p.CellSize))
}
