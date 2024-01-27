package main

import (
	"image"
	"image/color"
	"math"

	"github.com/mlange-42/arche/ecs"
)

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

// Colors for bee activities, as a resource
type Colors struct {
	Scout  color.RGBA
	Forage color.RGBA
	Return color.RGBA
	Waggle color.RGBA
	InHive color.RGBA
	Follow color.RGBA
}

// NewColors returns default bee activity colors.
func NewColors() Colors {
	return Colors{
		Scout:  color.RGBA{255, 255, 20, 255},
		Forage: color.RGBA{160, 160, 160, 160},
		Return: color.RGBA{0, 255, 255, 255},
		Waggle: color.RGBA{255, 50, 50, 255},
		InHive: color.RGBA{80, 80, 80, 255},
		Follow: color.RGBA{255, 255, 255, 255},
	}
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
