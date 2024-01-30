//go:build !js

package common

import "github.com/hajimehoshi/ebiten/v2"

type canvasHelper struct {
	width  int
	height int
}

func newCanvasHelper(width, height int) *canvasHelper {
	return &canvasHelper{
		width:  width,
		height: height,
	}
}

func (c *canvasHelper) isMouseInside() bool {
	x, y := ebiten.CursorPosition()
	return x >= 0 && y >= 0 && x < c.width && y < c.height
}
