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
		Octaves:   3,
		Falloff:   0.75,
		Cutoff:    0.45,
	})
	game.Model.AddSystem(&evolution.SysInitEntities{
		InitialCount:    50,
		ReleaseInterval: 60,
		ReleaseCount:    1,
		RandomGenes:     false,
	})

	game.Model.AddSystem(&evolution.SysGrowGrassLogistic{
		Interval: 4,
		BaseRate: 0.15,
	})
	game.Model.AddSystem(&evolution.SysGrazing{
		MaxUptake:    0.005,
		UptakeFactor: 1.0,
	})
	game.Model.AddSystem(&evolution.SysSearching{
		MaxSpeed: 0.5,
	})
	game.Model.AddSystem(&evolution.SysDecisions{})
	game.Model.AddSystem(&evolution.SysReproduction{
		MatingTrials:        10,
		MaxMatingDiff:       10,
		CrossProb:           0.15,
		MutationProbability: 0.25,
		MutationMagnitude:   0.025,
		AllowAsexual:        false,
		HatchRadius:         2.0,
	})
	game.Model.AddSystem(&evolution.SysMetabolism{
		RateGrazing:   0.002,
		RateSearching: 0.008,
	})
	game.Model.AddSystem(&evolution.SysMortality{
		MaxAge: 6000,
	})
	game.Model.AddSystem(&evolution.SysDisturbance{
		Interval:    600,
		Count:       1,
		MinRadius:   3,
		MaxRadius:   5,
		TargetValue: 0.01,
	})

	game.Model.AddUISystem(&common.UISysSimSpeed{
		InitialExponent: 1,
		MinExponent:     -2,
		MaxExponent:     6,
	})
	game.Model.AddUISystem(&evolution.UISysManagePause{})
	game.Model.AddUISystem(&evolution.UISysDrawGrass{})
	game.Model.AddUISystem(&evolution.UISysDrawEntities{})
	game.Model.AddUISystem(&evolution.UISysDrawScatter{
		Interval:    60,
		XIndex:      0,
		YIndex:      1,
		Width:       160,
		Height:      160,
		ImageOffset: evolution.Position{X: 680, Y: 0},
	})
	game.Model.AddUISystem(&evolution.UISysDrawScatter{
		Interval:    60,
		XIndex:      2,
		YIndex:      3,
		Width:       160,
		Height:      160,
		ImageOffset: evolution.Position{X: 680, Y: 160},
	})
	game.Model.AddUISystem(&evolution.UISysDrawScatter{
		Interval:    60,
		XIndex:      4,
		YIndex:      5,
		Width:       160,
		Height:      160,
		ImageOffset: evolution.Position{X: 680, Y: 320},
	})

	game.Initialize()
	if err := game.Run(); err != nil {
		log.Fatal(err)
	}
}
