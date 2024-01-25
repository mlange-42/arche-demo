package main

import (
	"image"

	"github.com/ByteArena/box2d"
)

// Body component
type Body struct {
	Body   *box2d.B2Body
	Radius float64
}

// BoxWorld resource
type BoxWorld struct {
	World *box2d.B2World
}

// Image resource
type Image struct {
	Image  *image.RGBA
	Width  int
	Height int
	Redraw func()
}
