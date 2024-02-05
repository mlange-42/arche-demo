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

	game.Model.AddSystem(&boids.SysInitBoids{Count: 50})

	game.Model.AddSystem(&boids.SysNeighbors{
		Neighbors: 6,
		Radius:    60,
		BuildStep: 4,
	})
	game.Model.AddSystem(&boids.SysMoveBoids{
		Speed:          1,
		UpdateInterval: 4,

		SeparationDist:   15,
		SeparationWeight: 0.03,
		CohesionWeight:   0.02,
		AlignmentWeight:  0.1,
		SpeedWeight:      0.1,

		WallDist:   150,
		WallWeight: 0.05,
	})

	game.Model.AddUISystem(&boids.UISysDrawBoids{})
	//game.Model.AddUISystem(&boids.UISysDrawBoidsLines{})
	/*game.Model.AddUISystem(&boids.UISysDrawBoid{
		Radius: 60,
	})*/

	game.Model.AddUISystem(&boids.UISysManagePause{})
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
