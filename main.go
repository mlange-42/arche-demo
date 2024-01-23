package main

import (
	"image/color"
	"syscall/js"
	"time"

	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/llgcode/draw2d/draw2dkit"
	"github.com/markfarnan/go-canvas/canvas"
	"github.com/mlange-42/arche-model/model"
	"github.com/mlange-42/arche/ecs"
)

var done chan struct{}

// Canvas resource
type Canvas struct {
	Canvas *canvas.Canvas2d
	Width  float64
	Height float64
}

var cvs Canvas
var mod *model.Model

// This specifies how long a delay between calls to 'render'.     To get Frame Rate,   1s / renderDelay
var renderDelay time.Duration = 20 * time.Millisecond

func main() {

	FrameRate := time.Second / renderDelay
	println("Hello Browser FPS:", FrameRate)
	//cvs, _ = canvas.NewCanvas2d(true)

	cvs.Canvas, _ = canvas.NewCanvas2d(false)
	cvs.Canvas.Create(int(js.Global().Get("innerWidth").Float()*0.9), int(js.Global().Get("innerHeight").Float()*0.9)) // Make Canvas 90% of window size.  For testing rendering canvas smaller than full windows

	cvs.Height = float64(cvs.Canvas.Height())
	cvs.Width = float64(cvs.Canvas.Width())

	mod = model.New()
	mod.FPS = 1000
	mod.TPS = 1000
	ecs.AddResource(&mod.World, &cvs)

	mod.Initialize()

	cvs.Canvas.Start(60, Render)
	//go doEvery(renderDelay, Render) // Kick off the Render function as go routine as it never returns
	<-done
}

// Helper function which calls the required func (in this case 'render') every time.Duration,  Call as a go-routine to prevent blocking, as this never returns
func doEvery(d time.Duration, f func(time.Time)) {
	for x := range time.Tick(d) {
		f(x)
	}
}

// Render is called from the 'requestAnnimationFrame' function.   It may also be called seperatly from a 'doEvery' function, if the user prefers drawing to be seperate from the annimationFrame callback
func Render(gc *draw2dimg.GraphicContext) bool {

	mod.Update()

	gc.SetFillColor(color.RGBA{0xff, 0xff, 0xff, 0xff})
	gc.Clear()
	// draws red 🔴 laser
	gc.SetFillColor(color.RGBA{0xff, 0x00, 0xff, 0xff})
	gc.SetStrokeColor(color.RGBA{0xff, 0x00, 0xff, 0xff})

	gc.BeginPath()
	//gc.ArcTo(gs.laserX, gs.laserY, gs.laserSize, gs.laserSize, 0, math.Pi*2)
	draw2dkit.Circle(gc, 100, 100, 20)
	gc.FillStroke()
	gc.Close()

	return true
}
