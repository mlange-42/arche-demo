package main

import (
	"image"
	"log"

	"github.com/mlange-42/arche-demo/common"
	"github.com/mlange-42/arche-demo/common/systems"
	"github.com/mlange-42/arche-demo/flocking"
	"github.com/mlange-42/arche-model/model"
	"github.com/mlange-42/arche/ecs"
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

	images := flocking.NewImages()
	ecs.AddResource(&game.Model.World, &images)

	game.Model.AddSystem(&flocking.SysInitBoids{Count: 300})

	game.Model.AddSystem(&flocking.SysNeighbors{
		Neighbors: 8,
		Radius:    50,
		BuildStep: 4,
	})
	game.Model.AddSystem(&flocking.SysMoveBoids{
		Speed:          1,
		UpdateInterval: 4,

		SeparationDist:  8,
		SeparationAngle: 3,
		CohesionAngle:   1.5,
		AlignmentAngle:  2,

		WallDist:  80,
		WallAngle: 12,

		MouseDist:  100,
		MouseAngle: 5,
	})

	//game.Model.AddUISystem(&boids.UISysDrawBoids{})
	game.Model.AddUISystem(&flocking.UISysDrawBoidsLines{})
	/*game.Model.AddUISystem(&boids.UISysDrawBoid{
		Radius: 60,
	})*/

	game.Model.AddUISystem(&flocking.UISysManagePause{})
	game.Model.AddUISystem(&systems.SimSpeed{
		InitialExponent: 0,
		MinExponent:     -2,
		MaxExponent:     2,
	})
	game.Model.AddUISystem(&systems.DrawInfo{
		Offset: image.Point{X: 800, Y: 0},
	})

	game.Initialize()
	if err := game.Run(); err != nil {
		log.Fatal(err)
	}
}
