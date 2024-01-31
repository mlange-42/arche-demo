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

	worldWidth  = 680
	worldHeight = 480
)

func main() {
	game := common.NewGame(
		model.New(), screenWidth, screenHeight,
	)

	grass := evolution.NewGrass(worldWidth, worldHeight, 4)
	ecs.AddResource(&game.Model.World, &grass)

	ecs.AddResource(&game.Model.World, &game.Screen)
	ecs.AddResource(&game.Model.World, &game.Mouse)

	game.Model.AddSystem(&evolution.SysInitGrass{
		Frequency: 0.05,
		Octaves:   3,
		Cutoff:    0.5,
	})
	game.Model.AddSystem(&evolution.SysInitEntities{
		Count: 10000,
	})

	game.Model.AddSystem(&evolution.SysGrowGrass{
		Interval: 4,
		BaseRate: 0.02,
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
		RateSearching: 0.005,
	})
	game.Model.AddSystem(&evolution.SysMortality{})
	game.Model.AddSystem(&evolution.SysReproduction{
		MatingTrials:  10,
		MaxMatingDiff: 15,
		CrossProb:     0.25,
		AllowAsexual:  false,
	})
	game.Model.AddSystem(&evolution.SysDisturbance{
		Interval:  10,
		Count:     1,
		MinRadius: 3,
		MaxRadius: 6,
	})

	game.Model.AddUISystem(&common.UISysSimSpeed{
		MinExponent: -2,
		MaxExponent: 4,
	})
	game.Model.AddUISystem(&evolution.UISysManagePause{})
	game.Model.AddUISystem(&evolution.UISysDrawGrass{})
	game.Model.AddUISystem(&evolution.UISysDrawEntities{})

	game.Initialize()
	if err := game.Run(); err != nil {
		log.Fatal(err)
	}
}
