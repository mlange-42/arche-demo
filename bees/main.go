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

	params := Params{
		MaxBeeSpeed: 1.0,
	}
	ecs.AddResource(&mod.World, &params)

	cvs, _ = common.NewCanvas("canvas-container", false)
	cvs.Create(int(math.Min(js.Global().Get("innerWidth").Float(), 880)), 480)

	image := Image{Image: cvs.Image, Width: cvs.Width, Height: cvs.Height, Redraw: cvs.Redraw}
	ecs.AddResource(&mod.World, &image)

	listener := MouseListener{}
	cvs.MouseListener = &listener
	ecs.AddResource(&mod.World, &listener)

	patches := NewPatches(image.Width, image.Height, 10)
	ecs.AddResource(&mod.World, &patches)

	colors := NewColors()
	ecs.AddResource(&mod.World, &colors)

	mod.AddSystem(&InitHives{Count: 5})
	mod.AddSystem(&InitBees{CountPerHive: 1000})

	mod.AddSystem(&ManagePatches{
		Count: 50,
	})
	mod.AddSystem(&SysHiveDecisions{
		ReleaseInterval:  8,
		ReleaseCount:     8,
		ScoutProbability: 0.1,
		DanceSamples:     2,
	})

	mod.AddSystem(&SysScouting{
		MaxRotation:  90,
		MaxScoutTime: 1500,
	})
	mod.AddSystem(&SysFollowing{
		MaxRotation:      45,
		ScoutProbability: 0.2,
	})
	mod.AddSystem(&SysForaging{
		MaxForagingTime: 120,
		MaxCollect:      0.001,
	})
	mod.AddSystem(&SysReturning{
		MaxRotation:         45,
		FleeDistance:        80,
		MaxDanceProbability: 0.5,
	})
	mod.AddSystem(&SysWaggleDance{
		MinDanceDuration: 60,
		MaxDanceDuration: 600,
	})
	mod.AddSystem(&SysFleeing{
		FleeDistance: 50,
	})

	mod.AddUISystem(&ManagePause{})
	mod.AddUISystem(&DrawPatches{})
	mod.AddUISystem(&DrawBees{})
	mod.AddUISystem(&DrawHives{})
	mod.AddUISystem(&SysRepaint{})

	println("Running the model")
	mod.Run()
}
