package main

import (
	"log"

	"github.com/mlange-42/arche-demo/common"
	"github.com/mlange-42/arche-demo/evolution"
	"github.com/mlange-42/arche-model/model"
	"github.com/mlange-42/arche/ecs"
)

const (
	screenWidth  = 880
	screenHeight = 480

	worldWidth  = 340
	worldHeight = 240
)

func main() {
	game := common.NewGame(
		model.New(), screenWidth, screenHeight,
	)

	grass := evolution.NewGrass(worldWidth, worldHeight, 4, 2)
	ecs.AddResource(&game.Model.World, &grass)

	ecs.AddResource(&game.Model.World, &game.Screen)
	ecs.AddResource(&game.Model.World, &game.Mouse)

	game.Model.AddSystem(&evolution.SysInitGrass{
		Frequency: 0.05,
		Octaves:   4,
		Cutoff:    0.45,
	})
	game.Model.AddSystem(&evolution.SysInitEntities{
		Count: 1000,
	})

	game.Model.AddSystem(&evolution.SysGrowGrass{
		Interval: 4,
		BaseRate: 0.04,
	})
	game.Model.AddSystem(&evolution.SysGrazing{
		MaxUptake: 0.01,
	})
	game.Model.AddSystem(&evolution.SysSearching{
		MaxSpeed: 0.5,
	})
	game.Model.AddSystem(&evolution.SysDecisions{})
	game.Model.AddSystem(&evolution.SysMetabolism{
		RateGrazing:   0.002,
		RateSearching: 0.01,
	})
	game.Model.AddSystem(&evolution.SysMortality{})
	game.Model.AddSystem(&evolution.SysReproduction{
		MatingTrials:      10,
		MaxMatingDiff:     15,
		CrossProb:         0.1,
		MutationMagnitude: 0.02,
		AllowAsexual:      true,
	})
	game.Model.AddSystem(&evolution.SysDisturbance{
		Interval:  60,
		Count:     1,
		MinRadius: 3,
		MaxRadius: 6,
	})

	game.Model.AddUISystem(&common.UISysSimSpeed{
		MinExponent: -2,
		MaxExponent: 5,
	})
	game.Model.AddUISystem(&evolution.UISysManagePause{})
	game.Model.AddUISystem(&evolution.UISysDrawGrass{})
	game.Model.AddUISystem(&evolution.UISysDrawEntities{})
	game.Model.AddUISystem(&evolution.UISysDrawScatter{
		Interval:    60,
		XIndex:      0,
		YIndex:      1,
		ImageOffset: evolution.Position{X: 680, Y: 0},
	})
	game.Model.AddUISystem(&evolution.UISysDrawScatter{
		Interval:    60,
		XIndex:      2,
		YIndex:      3,
		ImageOffset: evolution.Position{X: 680, Y: 200},
	})

	game.Initialize()
	if err := game.Run(); err != nil {
		log.Fatal(err)
	}
}
