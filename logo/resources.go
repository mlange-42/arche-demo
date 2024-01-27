package logo

import (
	"embed"
	"image"
)

// Logo is the embedded Ache logo.
//
//go:embed arche-logo-text.png
var Logo embed.FS

// Grid resource, holding the logo image data.
type Grid struct {
	Data   [][]bool
	Width  int
	Height int
}

// Image resource for drawing. Will be shown on an HTML5 canvas.
type Image struct {
	Image  *image.RGBA
	Width  int
	Height int
	Redraw func()
}
