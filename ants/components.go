package ants

import "github.com/mlange-42/arche/ecs"

// Position component
type Position struct {
	X float64
	Y float64
}

// Node component
type Node struct {
	NumEdges int
	InEdges  [8]ecs.Entity
	OutEdges [8]ecs.Entity
}

// Edge component
type Edge struct {
	From ecs.Entity
	To   ecs.Entity
}
