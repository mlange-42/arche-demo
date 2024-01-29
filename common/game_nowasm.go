//go:build !js

package common

type canvasHelper struct{}

func newCanvasHelper() canvasHelper {
	return canvasHelper{}
}

func (c *canvasHelper) isMouseInside() bool {
	return true
}
