package logo

import (
	"image/png"
	"log"

	"github.com/mlange-42/arche-demo/common"
	"github.com/mlange-42/arche-model/model"
	"github.com/mlange-42/arche/ecs"
)

const (
	screenWidth  = 880
	screenHeight = 480
)

func Run() {
	game := common.NewGame(
		model.New(), screenWidth, screenHeight,
	)

	grid, err := createImageResource()
	if err != nil {
		println("unable to load image: ", err.Error())
		panic(err)
	}
	ecs.AddResource(&game.Model.World, &grid)

	ecs.AddResource(&game.Model.World, &game.Screen)
	ecs.AddResource(&game.Model.World, &game.Mouse)

	game.Model.AddSystem(&SysInitEntities{})

	game.Model.AddSystem(&SysMoveEntities{
		MaxSpeed: 10,
		MaxAcc:   0.08, MaxAccFlee: 0.1,
		MinFleeDistance: 50,
		MaxFleeDistance: 200,
		Damp:            0.975})

	game.Model.AddUISystem(&UISysManagePause{})
	game.Model.AddUISystem(&UISysDrawEntities{})

	game.Initialize()
	if err := game.Run(); err != nil {
		log.Fatal(err)
	}
}

func createImageResource() (Grid, error) {
	f, err := Logo.Open("arche-logo-text.png")
	if err != nil {
		return Grid{}, err
	}
	defer f.Close()
	img, err := png.Decode(f)
	if err != nil {
		return Grid{}, err
	}
	w := img.Bounds().Dx()
	h := img.Bounds().Dy()
	data := make([][]bool, h)

	for i := 0; i < h; i++ {
		data[i] = make([]bool, w)
		for j := 0; j < w; j++ {
			r, _, _, _ := img.At(j, i).RGBA()
			data[i][j] = r > 32000
		}
	}

	return Grid{
		Data:   data,
		Width:  w,
		Height: h,
	}, nil
}
