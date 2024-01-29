package common

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

// Image resource for drawing. Will be shown on an HTML5 canvas.
type Image struct {
	Image  *image.RGBA
	Width  int
	Height int
	Redraw func()
}

// EbitenImage resource for drawing. Will be shown on an HTML5 canvas.
type EbitenImage struct {
	Image  *ebiten.Image
	Width  int
	Height int
}

// Mouse resource for events.
type Mouse struct {
	IsInside bool
	X        int
	Y        int
}
