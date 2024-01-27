package main

import (
	"github.com/mlange-42/arche-demo/ants"
	"github.com/mlange-42/arche-demo/common"
	"github.com/mlange-42/arche-model/model"
	"github.com/mlange-42/arche/ecs"
)

var cvs common.Canvas
var mod *model.Model

func main() {
	mod = model.New()
	mod.FPS = 60
	mod.TPS = 60

	cvs, _ = common.NewCanvas("canvas-container", 880, 480, true)

	image := common.Image{Image: cvs.Image(), Width: cvs.Width(), Height: cvs.Height(), Redraw: cvs.Redraw}
	ecs.AddResource(&mod.World, &image)

	grid := ants.NewPatches(image.Width, image.Height, 20)
	ecs.AddResource(&mod.World, &grid)

	listener := common.PauseMouseListener{}
	cvs.SetListener(&listener)
	ecs.AddResource(&mod.World, &listener)

	mod.AddSystem(&ants.InitGrid{})

	mod.AddSystem(&ants.SysDecay{
		Persistence: 0.99,
	})

	mod.AddUISystem(&ants.ManagePause{})
	mod.AddUISystem(&ants.DrawGrid{})

	println("Running the model")
	mod.Run()
}
