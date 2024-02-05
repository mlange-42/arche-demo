package boids

import (
	"math"

	"github.com/mlange-42/arche-demo/common"
	"github.com/mlange-42/arche/ecs"
)

const MaxNeighbors = 8

// Position component.
type Position struct {
	common.Vec2f
}

// Heading component.
type Heading struct {
	Angle float64
}

// Direction returns the unit vector corresponding to the heading.
func (h *Heading) Direction() (float32, float32) {
	a := float64(h.Angle)
	return float32(math.Cos(a)), float32(math.Sin(a))
}

// Neighbors holds an entity's neighbours
type Neighbors struct {
	Entities [MaxNeighbors]ecs.Entity
	Count    int
}
