package main

import (
	"fmt"
	"image"
	"syscall/js"
	"time"
)

// Canvas resource
type Canvas struct {
	Width  int
	Height int
	Image  *image.RGBA // The frame we actually draw on

	Mouse       Position
	MouseInside bool
	Paused      bool

	touchStart time.Time

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
}

// NewCanvas creates a new Canvas.
func NewCanvas(parentID string, create bool) (*Canvas, error) {
	var c Canvas

	c.window = js.Global()
	c.doc = c.window.Get("document")
	c.parent = c.doc.Call("getElementById", parentID)
	c.instructions = c.doc.Call("getElementById", "instructions")

	// If create, make a canvas that fills the windows
	if create {
		c.Create(int(c.window.Get("innerWidth").Int()), int(c.window.Get("innerHeight").Int()))
	}

	return &c, nil
}

// Create a new Canvas in the DOM, and append it to the Body.
// This also calls Set to create relevant shadow Buffer etc
func (c *Canvas) Create(width int, height int) {

	// Make the Canvas
	canvas := c.doc.Call("createElement", "canvas")

	canvas.Set("height", height)
	canvas.Set("width", width)
	c.parent.Call("appendChild", canvas)

	c.Set(canvas, width, height)
}

// Set up with an existing Canvas element which was obtained from JS
func (c *Canvas) Set(canvas js.Value, width int, height int) {
	c.canvas = canvas
	c.Height = height
	c.Width = width

	// Setup the 2D Drawing context
	c.ctx = c.canvas.Call("getContext", "2d")
	c.imgData = c.ctx.Call("createImageData", width, height) // Note Width, then Height
	c.Image = image.NewRGBA(image.Rect(0, 0, width, height))
	c.copybuff = js.Global().Get("Uint8Array").New(width * height * 4) // Static JS buffer for copying data out to JS. Defined once and re-used to save on un-needed allocations

	if c.isTouchDevice() {
		c.canvas.Set("ontouchstart", js.FuncOf(c.onTouchStart))
		c.canvas.Set("ontouchend", js.FuncOf(c.onTouchEnd))
		c.canvas.Set("ontouchcancel", js.FuncOf(c.onTouchEnd))
		c.canvas.Set("ontouchmove", js.FuncOf(c.onMouseMove))
	} else {
		c.canvas.Set("onmousemove", js.FuncOf(c.onMouseMove))
		c.canvas.Set("onmouseleave", js.FuncOf(c.onMouseLeave))
		c.canvas.Set("onmouseenter", js.FuncOf(c.onMouseEnter))
		c.canvas.Set("onclick", js.FuncOf(c.onClick))
	}
}

func (c *Canvas) isTouchDevice() bool {
	nav := c.window.Get("navigator")
	mxPts := nav.Get("maxTouchPoints")
	msMxPts := nav.Get("msMaxTouchPoints")
	return c.window.Call("hasOwnProperty", "ontouchstart").Bool() ||
		(!mxPts.IsUndefined() && mxPts.Int() > 0) ||
		(!msMxPts.IsUndefined() && msMxPts.Int() > 0)
}

func (c *Canvas) onMouseMove(this js.Value, args []js.Value) interface{} {
	evt := args[0]
	rect := c.canvas.Call("getBoundingClientRect")

	c.Mouse.X = float64(evt.Get("clientX").Int() - rect.Get("left").Int())
	c.Mouse.Y = float64(evt.Get("clientY").Int() - rect.Get("top").Int())

	c.instructions.Set("innerHTML", fmt.Sprintf("--> Moved to %.1f / %.1f", c.Mouse.X, c.Mouse.Y))
	return nil
}

func (c *Canvas) onMouseEnter(this js.Value, args []js.Value) interface{} {
	c.MouseInside = true
	c.instructions.Set("innerHTML", "Mouse entered")
	return nil
}

func (c *Canvas) onMouseLeave(this js.Value, args []js.Value) interface{} {
	c.MouseInside = false
	c.instructions.Set("innerHTML", "Mouse left")
	return nil
}

func (c *Canvas) onTouchStart(this js.Value, args []js.Value) interface{} {
	c.MouseInside = true
	c.touchStart = time.Now()
	c.instructions.Set("innerHTML", "Touch start")
	return nil
}

func (c *Canvas) onTouchEnd(this js.Value, args []js.Value) interface{} {
	c.MouseInside = false
	t := time.Now()
	if t.Sub(c.touchStart) < time.Second {
		c.instructions.Set("innerHTML", "Touch end --> Interpreted as click")
		c.Paused = !c.Paused
	} else {
		c.instructions.Set("innerHTML", "Touch end")
	}
	return nil
}

func (c *Canvas) onClick(this js.Value, args []js.Value) interface{} {
	c.Paused = !c.Paused
	c.instructions.Set("innerHTML", "Mouse clicked")
	return nil
}

// Redraw does the actuall copy over of the image data for the 'render' call.
func (c *Canvas) Redraw() {
	// TODO:  This currently does multiple data copies.   go image buffer -> JS Uint8Array,   Then JS Uint8Array -> ImageData,  then ImageData into the Canvas.
	// Would like to eliminate at least one of them, however currently CopyBytesToJS only supports Uint8Array  rather than the Uint8ClampedArray of ImageData.

	js.CopyBytesToJS(c.copybuff, c.Image.Pix)
	c.imgData.Get("data").Call("set", c.copybuff)
	c.ctx.Call("putImageData", c.imgData, 0, 0)
}
