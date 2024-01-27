//go:build !js

package common

import (
	"image"
)

// ImageCanvas implementation
type ImageCanvas struct {
	width  int
	height int
	image  *image.RGBA
}

// NewCanvas creates a new Canvas.
func NewCanvas(parentID string, width, height int, removeChildren bool) (Canvas, error) {
	var c ImageCanvas

	c.height = height
	c.width = width

	c.image = image.NewRGBA(image.Rect(0, 0, width, height))

	return &c, nil
}

// Image of the canvas
func (c *ImageCanvas) Image() *image.RGBA {
	return c.image
}

// Width of the canvas
func (c *ImageCanvas) Width() int {
	return c.width
}

// Height of the canvas
func (c *ImageCanvas) Height() int {
	return c.height
}

// SetListener of the canvas
func (c *ImageCanvas) SetListener(l MouseListener) {}

// Redraw does nothing in this canvas implementation.
func (c *ImageCanvas) Redraw() {}
