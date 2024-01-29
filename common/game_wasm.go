//go:build js

package common

import "syscall/js"

type canvasHelper struct {
	doc         js.Value
	canvas      js.Value
	mouseInside bool
}

func newCanvasHelper() *canvasHelper {
	doc := js.Global().Get("document")
	canvas := doc.Call("getElementsByTagName", "canvas").Index(0)

	helper := canvasHelper{
		doc:    doc,
		canvas: canvas,
	}

	canvas.Set("onmouseleave", js.FuncOf(helper.onMouseLeave))
	canvas.Set("onmouseenter", js.FuncOf(helper.onMouseEnter))

	return &helper
}

func (c *canvasHelper) isMouseInside() bool {
	return c.mouseInside
}

func (c *canvasHelper) onMouseEnter(this js.Value, args []js.Value) interface{} {
	c.mouseInside = true
	return nil
}

func (c *canvasHelper) onMouseLeave(this js.Value, args []js.Value) interface{} {
	c.mouseInside = false
	return nil
}
