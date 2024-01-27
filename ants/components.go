package ants

import "github.com/mlange-42/arche/ecs"

// Position component
type Position struct {
	X float64
	Y float64
}

// Coord component
type Coord struct {
	X int
	Y int
}

// Edge component
type Edge struct {
	From ecs.Entity
	To   ecs.Entity
}

// Node component
type Node struct {
	NumEdges int
	InEdges  [8]ecs.Entity
	OutEdges [8]ecs.Entity
}

// Add adds a pair of edges to the node.
func (n *Node) Add(in, out ecs.Entity) {
	n.InEdges[n.NumEdges] = in
	n.OutEdges[n.NumEdges] = out
	n.NumEdges++
}
