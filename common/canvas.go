package common

import "image"

// Canvas interface
type Canvas interface {
	Image() *image.RGBA
	Width() int
	Height() int
	SetListener(MouseListener)
	Redraw()
}
