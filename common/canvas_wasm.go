//go:build js

package common

import (
	"image"
	"syscall/js"
)

// HTMLCanvas implementation
type HTMLCanvas struct {
	width  int
	height int
	image  *image.RGBA // The frame we actually draw on

	mouseListener MouseListener

	// DOM properties
	window       js.Value
	doc          js.Value
	parent       js.Value
	instructions js.Value

	// Canvas properties
	canvas   js.Value
	ctx      js.Value
	imgData  js.Value
	copybuff js.Value

	repaint js.Func
}

// NewCanvas creates a new Canvas.
func NewCanvas(parentID string, width, height int, removeChildren bool) (Canvas, error) {
	var c HTMLCanvas

	c.window = js.Global()
	c.doc = c.window.Get("document")
	c.parent = c.doc.Call("getElementById", parentID)
	c.instructions = c.doc.Call("getElementById", "instructions")

	c.create(width, height, removeChildren)

	return &c, nil
}

// Image of the canvas
func (c *HTMLCanvas) Image() *image.RGBA {
	return c.image
}

// Width of the canvas
func (c *HTMLCanvas) Width() int {
	return c.width
}

// Height of the canvas
func (c *HTMLCanvas) Height() int {
	return c.height
}

// SetListener of the canvas
func (c *HTMLCanvas) SetListener(l MouseListener) {
	c.mouseListener = l
}

// create a new Canvas in the DOM, and append it to the Body.
// This also calls Set to create relevant shadow Buffer etc
func (c *HTMLCanvas) create(width int, height int, removeChildren bool) {
	// Remove other children
	if removeChildren {
		c.parent.Set("innerHTML", "")
	}

	// Make the Canvas
	canvas := c.doc.Call("createElement", "canvas")

	canvas.Set("height", height)
	canvas.Set("width", width)
	canvas.Set("id", "canvas")
	c.parent.Call("appendChild", canvas)

	c.set(canvas, width, height)
}

// set up with an existing Canvas element which was obtained from JS
func (c *HTMLCanvas) set(canvas js.Value, width int, height int) {
	c.canvas = canvas
	c.height = height
	c.width = width

	// Setup the 2D Drawing context
	c.ctx = c.canvas.Call("getContext", "2d")
	c.imgData = c.ctx.Call("createImageData", width, height) // Note Width, then Height
	c.image = image.NewRGBA(image.Rect(0, 0, width, height))
	c.copybuff = js.Global().Get("Uint8Array").New(width * height * 4) // Static JS buffer for copying data out to JS. Defined once and re-used to save on un-needed allocations

	c.canvas.Set("onmousemove", js.FuncOf(c.onMouseMove))
	c.canvas.Set("onmouseleave", js.FuncOf(c.onMouseLeave))
	c.canvas.Set("onmouseenter", js.FuncOf(c.onMouseEnter))
	c.canvas.Set("onclick", js.FuncOf(c.onClick))

	if c.isTouchDevice() {
		c.instructions.Set("innerHTML", "<s>"+c.instructions.Get("innerHTML").String()+"</s><br />Sorry, but this interactive simulation does not work with touch devices.")
	}

	c.repaint = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		//timestamp := args[0].Float()
		c.repaintCallback()
		return nil
	})
}

func (c *HTMLCanvas) isTouchDevice() bool {
	nav := c.window.Get("navigator")
	mxPts := nav.Get("maxTouchPoints")
	msMxPts := nav.Get("msMaxTouchPoints")
	return c.window.Call("hasOwnProperty", "ontouchstart").Bool() ||
		(!mxPts.IsUndefined() && mxPts.Int() > 0) ||
		(!msMxPts.IsUndefined() && msMxPts.Int() > 0)
}

func (c *HTMLCanvas) onMouseMove(this js.Value, args []js.Value) interface{} {
	if c.mouseListener != nil {
		c.mouseListener.MouseMove(c.getMousePosition(args[0]))
	}
	return nil
}

func (c *HTMLCanvas) getMousePosition(evt js.Value) MousePointer {
	rect := c.canvas.Call("getBoundingClientRect")

	scaleX := c.canvas.Get("width").Float() / rect.Get("width").Float()
	scaleY := c.canvas.Get("height").Float() / rect.Get("height").Float()

	return MousePointer{
		X: (float64(evt.Get("clientX").Int() - rect.Get("left").Int())) * scaleX,
		Y: (float64(evt.Get("clientY").Int() - rect.Get("top").Int())) * scaleY,
	}
}

func (c *HTMLCanvas) onMouseEnter(this js.Value, args []js.Value) interface{} {
	if c.mouseListener != nil {
		c.mouseListener.MouseEnter(c.getMousePosition(args[0]))
	}
	return nil
}

func (c *HTMLCanvas) onMouseLeave(this js.Value, args []js.Value) interface{} {
	if c.mouseListener != nil {
		c.mouseListener.MouseLeave(c.getMousePosition(args[0]))
	}
	return nil
}

func (c *HTMLCanvas) onClick(this js.Value, args []js.Value) interface{} {
	if c.mouseListener != nil {
		c.mouseListener.MouseClick(c.getMousePosition(args[0]))
	}
	return nil
}

// Redraw does the actual copy over of the image data for the 'render' call.
func (c *HTMLCanvas) Redraw() {
	js.CopyBytesToJS(c.copybuff, c.image.Pix)
	c.imgData.Get("data").Call("set", c.copybuff)

	js.Global().Call("requestAnimationFrame", c.repaint)
}

func (c *HTMLCanvas) repaintCallback() {
	c.ctx.Call("putImageData", c.imgData, 0, 0)
}

// RemoveElementByID removes an HTML element by ID.
func removeElementByID(doc js.Value, id string) {
	elem := doc.Call("getElementById", id)
	elem.Call("remove")
}
