package main

import (
	"image"

	"github.com/mlange-42/arche/ecs"
)

// Position component
type Position struct {
	X float64
	Y float64
}

// HomeHive component
type HomeHive struct {
	ecs.Relation
}

// Hive component
type Hive struct{}

// Image resource
type Image struct {
	Image  *image.RGBA
	Width  int
	Height int
	Redraw func()
}
