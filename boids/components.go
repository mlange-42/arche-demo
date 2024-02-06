package boids

import (
	"math"

	"github.com/mlange-42/arche-demo/common"
	"github.com/mlange-42/arche-demo/common/kd"
)

const MaxNeighbors = 8

// Position component.
type Position struct {
	common.Vec2f
}

/*
// Velocity component.
type Velocity struct {
	common.Vec2f
}
*/
// Heading component
type Heading struct {
	Angle float64
}

// Direction returns the unit vector corresponding to the heading.
func (h *Heading) Direction() common.Vec2f {
	a := h.Angle
	return common.Vec2f{X: math.Cos(a), Y: math.Sin(a)}
}

// Wrap brings the heading into rane [0, 2Pi).
func (h *Heading) Wrap() {
	h.Angle = math.Mod(h.Angle, 2*math.Pi)
	if h.Angle < 0 {
		h.Angle = 2*math.Pi - h.Angle
	}
}

type Rand256 struct {
	R uint8
}

// Neighbors holds an entity's neighbours
type Neighbors struct {
	Entities [MaxNeighbors]kd.EntityLocation
	Count    int
}
