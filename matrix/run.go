package matrix

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

	columnWidth = 10
	lineHeight  = 14
)

func Run() {
	game := common.NewGame(
		model.New(), screenWidth, screenHeight,
	)

	ecs.AddResource(&game.Model.World, &game.Screen)
	ecs.AddResource(&game.Model.World, &game.Mouse)
	letters := NewLetters()
	ecs.AddResource(&game.Model.World, &letters)
	grid := NewLetterGrid(screenWidth, screenHeight, columnWidth, lineHeight)
	ecs.AddResource(&game.Model.World, &grid)

	game.Model.AddSystem(&SysInitLetters{
		SpawnProb:       0.8,
		MinMoveInterval: 5,
		MaxMoveInterval: 7,
		MinGap:          60,
	})
	game.Model.AddSystem(&SysMoveLetters{})
	game.Model.AddSystem(&SysFadeLetters{
		FadeDuration: 120,
	})
	game.Model.AddSystem(&SysSwitchFaders{
		MinChangeInterval: 30,
		MaxChangeInterval: 180,
	})

	game.Model.AddUISystem(&UISysDrawLetters{})
	game.Model.AddUISystem(&systems.SimSpeed{
		InitialExponent: 0,
		MinExponent:     -2,
		MaxExponent:     4,
	})
	game.Model.AddUISystem(&systems.DrawInfo{
		Offset: image.Point{X: 800, Y: 0},
	})
	game.Model.AddUISystem(&UISysManagePause{})

	game.Initialize()
	if err := game.Run(); err != nil {
		log.Fatal(err)
	}
}
