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
}

// NewGrass creates a new Grass resource.
func NewGrass(width, height, cellsize int) Grass {
	return Grass{
		Grass:  common.NewGridExtent[float32](float64(width), float64(height), float64(cellsize)),
		Growth: common.NewGridExtent[float32](float64(width), float64(height), float64(cellsize)),
	}
}
