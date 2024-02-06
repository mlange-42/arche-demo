package ants

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

	ecs.AddResource(&game.Model.World, &game.Screen)
	ecs.AddResource(&game.Model.World, &game.Mouse)

	img := common.Image{
		Image:  image.NewRGBA(game.Screen.Image.Bounds()),
		Width:  game.Screen.Width,
		Height: game.Screen.Height,
	}
	ecs.AddResource(&game.Model.World, &img)

	grid := NewPatches(game.Screen.Width, game.Screen.Height, 10)
	ecs.AddResource(&game.Model.World, &grid)

	colors := NewColors()
	ecs.AddResource(&game.Model.World, &colors)

	game.Model.AddSystem(&SysInitGrid{})
	game.Model.AddSystem(&SysInitNest{
		AntsPerNest: 1000,
	})

	game.Model.AddSystem(&SysResources{
		Count: 32,
	})
	game.Model.AddSystem(&SysTraceDecay{
		Persistence: 0.99,
	})
	game.Model.AddSystem(&SysNestDecisions{
		ReleaseInterval:  8,
		ReleaseCount:     1,
		ScoutProbability: 0.1,
		ProbExponent:     0.6,
	})
	game.Model.AddSystem(&SysMoveAnts{
		MaxSpeed: 1.0,
	})
	game.Model.AddSystem(&SysScouting{
		MaxCollect:    0.001,
		TraceDecay:    0.95,
		MaxSearchTime: 1200,
	})
	game.Model.AddSystem(&SysForaging{
		MaxCollect:       0.001,
		ProbExponent:     1.0,
		RandomProb:       0.05,
		TraceDecay:       0.95,
		MaxSearchTime:    600,
		ScoutProbability: 0.05,
	})
	game.Model.AddSystem(&SysReturning{
		ProbExponent: 1.0,
		RandomProb:   0.05,
		TraceDecay:   0.95,
	})

	game.Model.AddUISystem(&UISysManagePause{})
	game.Model.AddUISystem(&UISysClearFrame{})
	game.Model.AddUISystem(&UISysDrawResources{})
	//game.Model.AddUISystem(&UISysDrawGrid{})
	game.Model.AddUISystem(&UISysDrawAnts{})
	game.Model.AddUISystem(&UISysDrawNest{})
	game.Model.AddUISystem(&UISysRepaint{})

	game.Model.AddUISystem(&systems.SimSpeed{
		InitialExponent: 0,
		MinExponent:     -2,
		MaxExponent:     6,
	})
	game.Model.AddUISystem(&systems.DrawInfo{
		Offset:     image.Point{X: 800, Y: 0},
		Components: generic.T1[Ant](),
	})

	game.Initialize()
	if err := game.Run(); err != nil {
		log.Fatal(err)
	}
}
