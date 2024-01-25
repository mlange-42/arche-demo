package main

import "image"

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

// Target component
type Target struct {
	X float64
	Y float64
}

// Grid resource
type Grid struct {
	Data   [][]bool
	Width  int
	Height int
}

// Image resource
type Image struct {
	Image  *image.RGBA
	Width  int
	Height int
	Redraw func()
}
