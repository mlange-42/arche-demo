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

	patches := NewPatches(image.Width, image.Height, 10)
	ecs.AddResource(&mod.World, &patches)

	mod.AddSystem(&InitHives{Count: 3})
	mod.AddSystem(&InitBees{CountPerHive: 1000})

	mod.AddSystem(&ManagePatches{
		Count: 25,
	})
	mod.AddSystem(&SysHiveDecisions{
		ReleaseInterval:  8,
		ReleaseCount:     4,
		ScoutProbability: 0.025,
		DanceSamples:     3,
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
		MaxCollect:      0.002,
	})
	mod.AddSystem(&SysReturning{
		MaxRotation: 45,
	})
	mod.AddSystem(&SysWaggleDance{
		MaxDanceDuration: 120,
	})

	mod.AddUISystem(&DrawHives{})

	println("Running the model")
	mod.Run()
}
