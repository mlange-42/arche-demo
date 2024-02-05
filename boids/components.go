package boids

import (
	"github.com/mlange-42/arche-demo/common"
	"github.com/mlange-42/arche-demo/common/kd"
)

const MaxNeighbors = 8

// Position component.
type Position struct {
	common.Vec2f
}

// Velocity component.
type Velocity struct {
	common.Vec2f
}

type Rand256 struct {
	R uint8
}

// Neighbors holds an entity's neighbours
type Neighbors struct {
	Entities [MaxNeighbors]kd.EntityLocation
	Count    int
}
