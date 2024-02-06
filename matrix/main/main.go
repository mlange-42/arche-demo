package main

import (
	"image"
	"log"

	"github.com/mlange-42/arche-demo/common"
	"github.com/mlange-42/arche-demo/common/systems"
	"github.com/mlange-42/arche-demo/matrix"
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

	game.Model.AddSystem(&matrix.SysInitLetters{Count: 1000})

	game.Model.AddUISystem(&matrix.UISysDrawLetters{})
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
