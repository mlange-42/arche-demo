package main

import (
	"image"
	"log"

	"github.com/mlange-42/arche-demo/bees"
	"github.com/mlange-42/arche-demo/boids"
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

	ecs.AddResource(&game.Model.World, &game.Screen)
	ecs.AddResource(&game.Model.World, &game.Mouse)

	img := common.Image{
		Image:  image.NewRGBA(game.Screen.Image.Bounds()),
		Width:  game.Screen.Width,
		Height: game.Screen.Height,
	}
	ecs.AddResource(&game.Model.World, &img)

	images := boids.NewImages()
	ecs.AddResource(&game.Model.World, &images)

	game.Model.AddSystem(&boids.SysInitBoids{Count: 500})

	game.Model.AddSystem(&boids.SysNeighbors{
		Neighbors: 8,
		Radius:    50,
		BuildStep: 8,
	})
	game.Model.AddSystem(&boids.SysMoveBoids{
		Speed:          0.5,
		UpdateInterval: 4,

		SeparationDist:   20,
		SeparationWeight: 0.02,
		CohesionWeight:   0.01,
		AlignmentWeight:  0.01,
		SpeedWeight:      0.025,
	})

	game.Model.AddUISystem(&boids.UISysDrawBoids{})

	game.Model.AddUISystem(&systems.SimSpeed{
		InitialExponent: 1,
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
