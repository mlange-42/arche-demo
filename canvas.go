package main

import (
	"image"
	"syscall/js"
)

// Canvas resource
type Canvas struct {
	Width  int
	Height int
	Image  *image.RGBA // The frame we actually draw on

	// DOM properties
	window js.Value
	doc    js.Value
	body   js.Value

	// Canvas properties
	canvas   js.Value
	ctx      js.Value
	imgData  js.Value
	copybuff js.Value
}

// NewCanvas creates a new Canvas.
func NewCanvas(create bool) (*Canvas, error) {
	var c Canvas

	c.window = js.Global()
	c.doc = c.window.Get("document")
	c.body = c.doc.Get("body")

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
	c.body.Call("appendChild", canvas)

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

}

// Redraw does the actuall copy over of the image data for the 'render' call.
func (c *Canvas) Redraw() {
	// TODO:  This currently does multiple data copies.   go image buffer -> JS Uint8Array,   Then JS Uint8Array -> ImageData,  then ImageData into the Canvas.
	// Would like to eliminate at least one of them, however currently CopyBytesToJS only supports Uint8Array  rather than the Uint8ClampedArray of ImageData.

	js.CopyBytesToJS(c.copybuff, c.Image.Pix)
	c.imgData.Get("data").Call("set", c.copybuff)
	c.ctx.Call("putImageData", c.imgData, 0, 0)
}
