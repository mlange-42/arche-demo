package main

import (
	"math"
	"syscall/js"

	"github.com/mlange-42/arche-demo/common"
	"github.com/mlange-42/arche-model/model"
	"github.com/mlange-42/arche/ecs"
)

var cvs *common.Canvas
var mod *model.Model

func main() {
	mod = model.New()
	mod.FPS = 60
	mod.TPS = 60

	common.RemoveElementByID("loading")

	cvs, _ = common.NewCanvas("canvas-container", false)
	cvs.Create(int(math.Min(js.Global().Get("innerWidth").Float(), 880)), 480)

	image := Image{Image: cvs.Image, Width: cvs.Width, Height: cvs.Height, Redraw: cvs.Redraw}
	ecs.AddResource(&mod.World, &image)

	println("Running the model")
	mod.Run()
}
