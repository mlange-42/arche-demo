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

	grid := ants.NewPatches(image.Width, image.Height, 10)
	ecs.AddResource(&mod.World, &grid)

	listener := common.PauseMouseListener{}
	cvs.SetListener(&listener)
	ecs.AddResource(&mod.World, &listener)

	colors := ants.NewColors()
	ecs.AddResource(&mod.World, &colors)

	mod.AddSystem(&ants.InitGrid{})
	mod.AddSystem(&ants.InitNest{
		AntsPerNest: 1000,
	})

	mod.AddSystem(&ants.SysResources{
		Count: 24,
	})
	mod.AddSystem(&ants.SysDecay{
		Persistence: 0.99,
	})
	mod.AddSystem(&ants.SysNestDecisions{
		ReleaseInterval:  8,
		ReleaseCount:     1,
		ScoutProbability: 0.1,
		ProbExponent:     0.6,
	})
	mod.AddSystem(&ants.SysMoveAnts{
		MaxSpeed: 1.0,
	})
	mod.AddSystem(&ants.SysScouting{
		MaxCollect:    0.001,
		TraceDecay:    0.95,
		MaxSearchTime: 600,
	})
	mod.AddSystem(&ants.SysForaging{
		MaxCollect:    0.001,
		ProbExponent:  1.0,
		RandomProb:    0.05,
		TraceDecay:    0.95,
		MaxSearchTime: 300,
	})
	mod.AddSystem(&ants.SysReturning{
		ProbExponent: 1.0,
		RandomProb:   0.05,
		TraceDecay:   0.95,
	})

	mod.AddUISystem(&ants.ManagePause{})
	mod.AddUISystem(&ants.SysClearFrame{})
	mod.AddUISystem(&ants.DrawResources{})
	mod.AddUISystem(&ants.DrawAnts{})
	mod.AddUISystem(&ants.DrawNest{})
	mod.AddUISystem(&ants.SysRepaint{})

	println("Running the model")
	mod.Run()
}
