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

// Image resource for drawing. Will be shown on an HTML5 canvas.
type Image struct {
	Image  *image.RGBA
	Width  int
	Height int
	Redraw func()
}

// Images resource, containing images for use in drawing Box2D bodies.
type Images struct {
	Circle *image.RGBA
}
