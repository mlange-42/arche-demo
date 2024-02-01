package main

import (
	"image"
	"log"

	"github.com/mlange-42/arche-demo/ants"
	"github.com/mlange-42/arche-demo/common"
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

	ecs.AddResource(&game.Model.World, &game.Screen)
	ecs.AddResource(&game.Model.World, &game.Mouse)

	img := common.Image{
		Image:  image.NewRGBA(game.Screen.Image.Bounds()),
		Width:  game.Screen.Width,
		Height: game.Screen.Height,
	}
	ecs.AddResource(&game.Model.World, &img)

	grid := ants.NewPatches(game.Screen.Width, game.Screen.Height, 10)
	ecs.AddResource(&game.Model.World, &grid)

	colors := ants.NewColors()
	ecs.AddResource(&game.Model.World, &colors)

	game.Model.AddSystem(&ants.SysInitGrid{})
	game.Model.AddSystem(&ants.SysInitNest{
		AntsPerNest: 1000,
	})

	game.Model.AddSystem(&ants.SysResources{
		Count: 32,
	})
	game.Model.AddSystem(&ants.SysDecay{
		Persistence: 0.99,
	})
	game.Model.AddSystem(&ants.SysNestDecisions{
		ReleaseInterval:  8,
		ReleaseCount:     1,
		ScoutProbability: 0.1,
		ProbExponent:     0.6,
	})
	game.Model.AddSystem(&ants.SysMoveAnts{
		MaxSpeed: 1.0,
	})
	game.Model.AddSystem(&ants.SysScouting{
		MaxCollect:    0.001,
		TraceDecay:    0.95,
		MaxSearchTime: 1200,
	})
	game.Model.AddSystem(&ants.SysForaging{
		MaxCollect:       0.001,
		ProbExponent:     1.0,
		RandomProb:       0.05,
		TraceDecay:       0.95,
		MaxSearchTime:    300,
		ScoutProbability: 0.05,
	})
	game.Model.AddSystem(&ants.SysReturning{
		ProbExponent: 1.0,
		RandomProb:   0.05,
		TraceDecay:   0.95,
	})

	game.Model.AddUISystem(&ants.UISysManagePause{})
	game.Model.AddUISystem(&ants.UISysClearFrame{})
	game.Model.AddUISystem(&ants.UISysDrawResources{})
	//game.Model.AddUISystem(&ants.UISysDrawGrid{})
	game.Model.AddUISystem(&ants.UISysDrawAnts{})
	game.Model.AddUISystem(&ants.UISysDrawNest{})
	game.Model.AddUISystem(&ants.UISysRepaint{})

	game.Model.AddUISystem(&common.UISysSimSpeed{
		InitialExponent: 1,
		MinExponent:     -2,
		MaxExponent:     6,
	})
	game.Model.AddUISystem(&common.UISysDrawInfo{
		Offset:     image.Point{X: 800, Y: 0},
		Components: generic.T1[ants.Ant](),
	})

	game.Initialize()
	if err := game.Run(); err != nil {
		log.Fatal(err)
	}
}
