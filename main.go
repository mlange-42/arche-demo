package main

import (
	"syscall/js"
	"time"

	"github.com/mlange-42/arche-model/model"
	"github.com/mlange-42/arche/ecs"
)

var done chan struct{}

var cvs *Canvas
var mod *model.Model

// This specifies how long a delay between calls to 'render'.     To get Frame Rate,   1s / renderDelay
var renderDelay time.Duration = 20 * time.Millisecond

func main() {
	cvs, _ = NewCanvas(false)
	cvs.Create(int(js.Global().Get("innerWidth").Float()*0.9), int(js.Global().Get("innerHeight").Float()*0.9)) // Make Canvas 90% of window size.  For testing rendering canvas smaller than full windows

	mod = model.New()
	mod.FPS = 60
	mod.TPS = 60
	ecs.AddResource(&mod.World, cvs)

	mod.AddSystem(&InitEntities{Count: 10000})
	mod.AddSystem(&MoveEntities{})
	mod.AddUISystem(&DrawEntities{})

	println("Running the model")
	mod.Run()
}
