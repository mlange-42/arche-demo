package evolution

import (
	"image"
	"log"

	"github.com/mlange-42/arche-demo/common"
	"github.com/mlange-42/arche-demo/common/systems"
	"github.com/mlange-42/arche-model/model"
	"github.com/mlange-42/arche/ecs"
)

const (
	screenWidth  = 880
	screenHeight = 480

	worldWidth  = 640
	worldHeight = 480

	scale = 2
)

func Run() {
	game := common.NewGame(
		model.New(), screenWidth, screenHeight,
	)

	grass := NewGrass(worldWidth/scale, worldHeight/scale, 4, scale)
	ecs.AddResource(&game.Model.World, &grass)

	selection := MouseSelection{}
	ecs.AddResource(&game.Model.World, &selection)

	ecs.AddResource(&game.Model.World, &game.Screen)
	ecs.AddResource(&game.Model.World, &game.Mouse)

	game.Model.AddSystem(&SysInitGrass{
		Frequency: 0.03,
		Octaves:   3,
		Falloff:   0.75,
		Cutoff:    0.4,
	})
	game.Model.AddSystem(&SysInitEntities{
		InitialBatches:  1000,
		ReleaseInterval: 240,
		ReleaseBatches:  1,
		BatchSize:       5,
		RandomGenes:     true,
	})

	game.Model.AddSystem(&SysGrowGrassLogistic{
		Interval: 4,
		BaseRate: 0.15,
	})
	game.Model.AddSystem(&SysGrazing{
		MaxUptake:    0.005,
		UptakeFactor: 1.0,
	})
	game.Model.AddSystem(&SysSearching{
		MaxSpeed: 0.5,
	})
	game.Model.AddSystem(&SysDecisions{})
	game.Model.AddSystem(&SysReproduction{
		MatingTrials:        10,
		MaxMatingDist:       60,
		MaxMatingDiff:       15,
		CrossProb:           0.333,
		MutationProbability: 0.2,
		MutationMagnitude:   0.01,
		AllowAsexual:        false,
		HatchRadius:         2.0,
	})
	game.Model.AddSystem(&SysMetabolism{
		RateGrazing:   0.002,
		RateSearching: 0.008,
	})
	game.Model.AddSystem(&SysMortality{
		MaxAge: 6000,
	})
	game.Model.AddSystem(&SysDisturbance{
		Interval:    600,
		Count:       0,
		MinRadius:   3,
		MaxRadius:   5,
		TargetValue: 0.01,
	})

	game.Model.AddUISystem(&systems.SimSpeed{
		InitialExponent: 1,
		MinExponent:     -2,
		MaxExponent:     6,
	})
	game.Model.AddUISystem(&UISysManagePause{})

	game.Model.AddUISystem(&UISysDrawGrass{})
	game.Model.AddUISystem(&UISysDrawEntities{})
	game.Model.AddUISystem(&UISysDrawScatter{
		Interval:       30,
		IntervalOffset: 0,
		XIndex:         0,
		YIndex:         1,
		Width:          160,
		Height:         160,
		ImageOffset:    image.Point{X: 640, Y: 0},
	})
	game.Model.AddUISystem(&UISysDrawScatter{
		Interval:       30,
		IntervalOffset: 10,
		XIndex:         2,
		YIndex:         3,
		Width:          160,
		Height:         160,
		ImageOffset:    image.Point{X: 640, Y: 160},
	})
	game.Model.AddUISystem(&UISysDrawScatter{
		Interval:       30,
		IntervalOffset: 20,
		XIndex:         4,
		YIndex:         5,
		Width:          160,
		Height:         160,
		ImageOffset:    image.Point{X: 640, Y: 320},
	})
	game.Model.AddUISystem(&systems.DrawInfo{
		Offset: image.Point{X: 800, Y: 0},
	})

	game.Initialize()
	if err := game.Run(); err != nil {
		log.Fatal(err)
	}
}
