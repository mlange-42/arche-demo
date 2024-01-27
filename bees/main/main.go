package main

import (
	"github.com/mlange-42/arche-demo/bees"
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

	params := bees.Params{
		MaxBeeSpeed: 1.0,
	}
	ecs.AddResource(&mod.World, &params)

	cvs, _ = common.NewCanvas("canvas-container", 880, 480, true)

	image := common.Image{Image: cvs.Image, Width: cvs.Width, Height: cvs.Height, Redraw: cvs.Redraw}
	ecs.AddResource(&mod.World, &image)

	listener := common.PauseMouseListener{}
	cvs.MouseListener = &listener
	ecs.AddResource(&mod.World, &listener)

	patches := bees.NewPatches(image.Width, image.Height, 10)
	ecs.AddResource(&mod.World, &patches)

	colors := bees.NewColors()
	ecs.AddResource(&mod.World, &colors)

	mod.AddSystem(&bees.InitHives{Count: 2})
	mod.AddSystem(&bees.InitBees{CountPerHive: 1000})

	mod.AddSystem(&bees.ManagePatches{
		Count: 12,
	})
	mod.AddSystem(&bees.SysHiveDecisions{
		ReleaseInterval:  8,
		ReleaseCount:     8,
		ScoutProbability: 0.1,
		DanceSamples:     2,
	})

	mod.AddSystem(&bees.SysScouting{
		MaxRotation:  90,
		MaxScoutTime: 1500,
	})
	mod.AddSystem(&bees.SysFollowing{
		MaxRotation:      45,
		ScoutProbability: 0.2,
	})
	mod.AddSystem(&bees.SysForaging{
		MaxForagingTime: 120,
		MaxCollect:      0.001,
	})
	mod.AddSystem(&bees.SysReturning{
		MaxRotation:         45,
		FleeDistance:        80,
		MaxDanceProbability: 0.5,
	})
	mod.AddSystem(&bees.SysWaggleDance{
		MinDanceDuration: 60,
		MaxDanceDuration: 600,
	})
	mod.AddSystem(&bees.SysFleeing{
		FleeDistance: 50,
	})

	mod.AddUISystem(&bees.ManagePause{})
	mod.AddUISystem(&bees.DrawPatches{})
	mod.AddUISystem(&bees.DrawBees{})
	mod.AddUISystem(&bees.DrawHives{})
	mod.AddUISystem(&bees.SysRepaint{})

	mod.Run()
}
