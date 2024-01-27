package boids

import "github.com/mlange-42/arche/ecs"

// Position component
type Position struct {
	X float64
	Y float64
}

// Velocity component
type Velocity struct {
	X float64
	Y float64
}

// Cell resource.
type Cell struct {
	X int
	Y int
}

// CurrentCell resource.
type CurrentCell struct {
	ecs.Relation
}
