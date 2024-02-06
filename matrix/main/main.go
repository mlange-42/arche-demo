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

	columnWidth = 10
	lineHeight  = 14
)

func main() {
	game := common.NewGame(
		model.New(), screenWidth, screenHeight,
	)

	ecs.AddResource(&game.Model.World, &game.Screen)
	ecs.AddResource(&game.Model.World, &game.Mouse)
	letters := matrix.NewLetters()
	ecs.AddResource(&game.Model.World, &letters)
	grid := matrix.NewLetterGrid(screenWidth, screenHeight, columnWidth, lineHeight)
	ecs.AddResource(&game.Model.World, &grid)

	gridManager := matrix.NewGridManager(&game.Model.World)
	game.Model.World.SetListener(&gridManager)

	game.Model.AddSystem(&matrix.SysInitLetters{
		SpawnProb:       0.8,
		MinMoveInterval: 5,
		MaxMoveInterval: 7,
		MinGap:          60,
	})
	game.Model.AddSystem(&matrix.SysMoveLetters{})
	game.Model.AddSystem(&matrix.SysFadeLetters{
		FadeDuration: 120,
	})

	game.Model.AddUISystem(&matrix.UISysDrawLetters{})
	game.Model.AddUISystem(&systems.SimSpeed{
		InitialExponent: 0,
		MinExponent:     -2,
		MaxExponent:     4,
	})
	game.Model.AddUISystem(&systems.DrawInfo{
		Offset: image.Point{X: 800, Y: 0},
	})
	game.Model.AddUISystem(&matrix.UISysManagePause{})

	game.Initialize()
	if err := game.Run(); err != nil {
		log.Fatal(err)
	}
}
