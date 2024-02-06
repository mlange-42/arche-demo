package main

import (
	"image"
	"log"

	"github.com/mlange-42/arche-demo/bees"
	"github.com/mlange-42/arche-demo/common"
	"github.com/mlange-42/arche-demo/common/systems"
	"github.com/mlange-42/arche-model/model"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
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

	img := common.Image{
		Image:  image.NewRGBA(game.Screen.Image.Bounds()),
		Width:  game.Screen.Width,
		Height: game.Screen.Height,
	}
	ecs.AddResource(&game.Model.World, &img)

	patches := bees.NewPatches(game.Screen.Width, game.Screen.Height, 10)
	ecs.AddResource(&game.Model.World, &patches)

	colors := bees.NewColors()
	ecs.AddResource(&game.Model.World, &colors)

	game.Model.AddSystem(&bees.SysInitHives{Count: 2})
	game.Model.AddSystem(&bees.SysInitBees{CountPerHive: 1000})

	game.Model.AddSystem(&bees.SysManagePatches{
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

	game.Model.AddUISystem(&bees.UISysManagePause{})
	game.Model.AddUISystem(&bees.UISysClearFrame{})
	game.Model.AddUISystem(&bees.UISysDrawPatches{})
	game.Model.AddUISystem(&bees.UISysDrawBees{})
	game.Model.AddUISystem(&bees.UISysDrawHives{})
	game.Model.AddUISystem(&bees.UISysRepaint{})

	game.Model.AddUISystem(&systems.SimSpeed{
		InitialExponent: 0,
		MinExponent:     -2,
		MaxExponent:     6,
	})
	game.Model.AddUISystem(&systems.DrawInfo{
		Offset:     image.Point{X: 800, Y: 0},
		Components: generic.T1[bees.HomeHive](),
	})

	game.Initialize()
	if err := game.Run(); err != nil {
		log.Fatal(err)
	}
}
