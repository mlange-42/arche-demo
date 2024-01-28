package ants

import "github.com/mlange-42/arche/ecs"

// Position component
type Position struct {
	X float64
	Y float64
}

// Edge component
type Edge struct {
	From ecs.Entity
	To   ecs.Entity
}

// EdgeGeometry component
type EdgeGeometry struct {
	From   Position
	Dir    Position
	Length float64
}

// Node component
type Node struct {
	NumEdges int
	InEdges  [8]ecs.Entity
	OutEdges [8]ecs.Entity
}

// NodeResource component
type NodeResource struct {
	Resource float64
}

// Add adds a pair of edges to the node.
func (n *Node) Add(in, out ecs.Entity) {
	n.InEdges[n.NumEdges] = in
	n.OutEdges[n.NumEdges] = out
	n.NumEdges++
}

// Trace component
type Trace struct {
	FromNest     float64
	FromResource float64
}

// Ant label component
type Ant struct{}

// AntEdge component
type AntEdge struct {
	Edge   ecs.Entity
	From   Position
	Dir    Position
	Length float64
	Pos    float64
}

// Update the AntEdge to the given edge
func (e *AntEdge) Update(edge ecs.Entity, geom *EdgeGeometry) {
	e.Edge = edge
	e.From = geom.From
	e.Dir = geom.Dir
	e.Length = geom.Length
	e.Pos = 0
}

// UpdatePos updates the argument to the ant's position along the edge
func (e *AntEdge) UpdatePos(pos *Position) {
	pos.X = e.From.X + e.Dir.X*e.Pos
	pos.Y = e.From.Y + e.Dir.Y*e.Pos
}

// ActScout component
type ActScout struct {
	Start int64
}

// ActForage component
type ActForage struct{}

// ActReturn component
type ActReturn struct {
	Load float64
}

// ActInNest component
type ActInNest struct{}
