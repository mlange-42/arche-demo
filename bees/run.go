package bees

import (
	"image"
	"log"

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

func Run() {
	game := common.NewGame(
		model.New(), screenWidth, screenHeight,
	)

	params := Params{
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

	patches := NewPatches(game.Screen.Width, game.Screen.Height, 10)
	ecs.AddResource(&game.Model.World, &patches)

	colors := NewColors()
	ecs.AddResource(&game.Model.World, &colors)

	game.Model.AddSystem(&SysInitHives{Count: 2})
	game.Model.AddSystem(&SysInitBees{CountPerHive: 1000})

	game.Model.AddSystem(&SysManagePatches{
		Count: 12,
	})
	game.Model.AddSystem(&SysHiveDecisions{
		ReleaseInterval:  8,
		ReleaseCount:     8,
		ScoutProbability: 0.1,
		DanceSamples:     2,
	})

	game.Model.AddSystem(&SysScouting{
		MaxRotation:  90,
		MaxScoutTime: 1500,
	})
	game.Model.AddSystem(&SysFollowing{
		MaxRotation:      45,
		ScoutProbability: 0.2,
	})
	game.Model.AddSystem(&SysForaging{
		MaxForagingTime: 120,
		MaxCollect:      0.001,
	})
	game.Model.AddSystem(&SysReturning{
		MaxRotation:         45,
		FleeDistance:        80,
		MaxDanceProbability: 0.5,
	})
	game.Model.AddSystem(&SysWaggleDance{
		MinDanceDuration: 60,
		MaxDanceDuration: 600,
	})
	game.Model.AddSystem(&SysFleeing{
		FleeDistance: 50,
	})

	game.Model.AddUISystem(&UISysManagePause{})
	game.Model.AddUISystem(&UISysClearFrame{})
	game.Model.AddUISystem(&UISysDrawPatches{})
	game.Model.AddUISystem(&UISysDrawBees{})
	game.Model.AddUISystem(&UISysDrawHives{})
	game.Model.AddUISystem(&UISysRepaint{})

	game.Model.AddUISystem(&systems.SimSpeed{
		InitialExponent: 0,
		MinExponent:     -2,
		MaxExponent:     6,
	})
	game.Model.AddUISystem(&systems.DrawInfo{
		Offset:     image.Point{X: 800, Y: 0},
		Components: generic.T1[HomeHive](),
	})

	game.Initialize()
	if err := game.Run(); err != nil {
		log.Fatal(err)
	}
}
