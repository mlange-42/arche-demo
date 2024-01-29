package main

import (
	"image"
	"log"

	"github.com/mlange-42/arche-demo/bees"
	"github.com/mlange-42/arche-demo/common"
	"github.com/mlange-42/arche-model/model"
	"github.com/mlange-42/arche/ecs"
)

const (
	screenWidth  = 880
	screenHeight = 480
)

func main() {
	game := common.NewGame(
		model.New(), screenWidth, screenHeight,
	)

	params := bees.Params{
		MaxBeeSpeed: 1.0,
	}
	ecs.AddResource(&game.Model.World, &params)
	ecs.AddResource(&game.Model.World, &game.Screen)
	ecs.AddResource(&game.Model.World, &game.Mouse)

	image := common.Image{
		Image:  image.NewRGBA(game.Screen.Image.Bounds()),
		Width:  game.Screen.Width,
		Height: game.Screen.Height,
	}
	ecs.AddResource(&game.Model.World, &image)

	patches := bees.NewPatches(game.Screen.Width, game.Screen.Height, 10)
	ecs.AddResource(&game.Model.World, &patches)

	colors := bees.NewColors()
	ecs.AddResource(&game.Model.World, &colors)

	game.Model.AddSystem(&bees.InitHives{Count: 2})
	game.Model.AddSystem(&bees.InitBees{CountPerHive: 1000})

	game.Model.AddSystem(&bees.ManagePatches{
		Count: 12,
	})
	game.Model.AddSystem(&bees.SysHiveDecisions{
		ReleaseInterval:  8,
		ReleaseCount:     8,
		ScoutProbability: 0.1,
		DanceSamples:     2,
	})

	game.Model.AddSystem(&bees.SysScouting{
		MaxRotation:  90,
		MaxScoutTime: 1500,
	})
	game.Model.AddSystem(&bees.SysFollowing{
		MaxRotation:      45,
		ScoutProbability: 0.2,
	})
	game.Model.AddSystem(&bees.SysForaging{
		MaxForagingTime: 120,
		MaxCollect:      0.001,
	})
	game.Model.AddSystem(&bees.SysReturning{
		MaxRotation:         45,
		FleeDistance:        80,
		MaxDanceProbability: 0.5,
	})
	game.Model.AddSystem(&bees.SysWaggleDance{
		MinDanceDuration: 60,
		MaxDanceDuration: 600,
	})
	game.Model.AddSystem(&bees.SysFleeing{
		FleeDistance: 50,
	})

	game.Model.AddUISystem(&bees.ManagePause{})
	game.Model.AddUISystem(&bees.SysClearFrame{})
	game.Model.AddUISystem(&bees.DrawPatches{})
	game.Model.AddUISystem(&bees.DrawBees{})
	game.Model.AddUISystem(&bees.DrawHives{})
	game.Model.AddUISystem(&bees.SysRepaint{})

	game.Initialize()
	if err := game.Run(); err != nil {
		log.Fatal(err)
	}
}
