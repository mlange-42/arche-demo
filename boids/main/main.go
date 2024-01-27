package main

import (
	"github.com/mlange-42/arche-demo/boids"
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

	cvs, _ = common.NewCanvas("canvas-container", 880, 480, true)

	image := common.Image{Image: cvs.Image, Width: cvs.Width, Height: cvs.Height, Redraw: cvs.Redraw}
	ecs.AddResource(&mod.World, &image)

	listener := common.PauseMouseListener{}
	cvs.MouseListener = &listener
	ecs.AddResource(&mod.World, &listener)

	mod.AddSystem(&boids.InitEntities{
		Count: 250,
	})

	mod.AddSystem(&boids.MoveEntities{})

	mod.AddUISystem(&boids.ManagePause{})

	mod.AddUISystem(&boids.DrawEntities{})

	println("Running the model")
	mod.Run()
}
