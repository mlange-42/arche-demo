package evolution

import (
	"github.com/mlange-42/arche-demo/common"
)

// GrasserCount resource.
type GrasserCount = common.Grid[int32]

// Grass resource.
type Grass struct {
	Grass  common.Grid[float32]
	Growth common.Grid[float32]
	Scale  int
}

// NewGrass creates a new Grass resource.
func NewGrass(width, height, cellsize int, scale int) Grass {
	return Grass{
		Grass:  common.NewGridExtent[float32](float64(width), float64(height), float64(cellsize)),
		Growth: common.NewGridExtent[float32](float64(width), float64(height), float64(cellsize)),
		Scale:  scale,
	}
}

// MouseSelection resource.
type MouseSelection struct {
	Selections []*SelectionEntry
}

// SelectionEntry for [MouseSelection].
type SelectionEntry struct {
	XIndex   int
	YIndex   int
	Position Position
	Radius   float32
	Active   bool
}
