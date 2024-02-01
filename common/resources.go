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
}

// EbitenImage resource for drawing. Will be shown on an HTML5 canvas.
type EbitenImage struct {
	Image  *ebiten.Image
	Width  int
	Height int
}

// Mouse resource for events.
type Mouse struct {
	image.Point
	IsInside bool
}

// SimulationSpeed resource.
// The [Game] will adapt the simulation speed if the resource is present.
type SimulationSpeed struct {
	Exponent int
}
