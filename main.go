package main

import (
	"embed"
	"image/png"
	"math"
	"syscall/js"

	"github.com/mlange-42/arche-model/model"
	"github.com/mlange-42/arche/ecs"
)

var cvs *Canvas
var mod *model.Model

//go:embed assets/arche-logo-text.png
var logo embed.FS

func main() {
	mod = model.New()
	mod.FPS = 60
	mod.TPS = 60

	grid, err := createImageResource()
	if err != nil {
		println("unable to load image: ", err.Error())
		panic(err)
	}
	ecs.AddResource(&mod.World, &grid)

	cvs, _ = NewCanvas("canvas", false)
	cvs.Create(int(math.Min(js.Global().Get("innerWidth").Float(), 960)), grid.Height*2)
	ecs.AddResource(&mod.World, cvs)

	mod.AddSystem(&InitEntities{})

	mod.AddSystem(&MoveEntities{
		MaxSpeed: 10,
		MaxAcc:   0.08, MaxAccFlee: 0.1,
		MinFleeDistance: 50,
		MaxFleeDistance: 200,
		Damp:            0.975})

	mod.AddUISystem(&ManagePause{})

	mod.AddUISystem(&DrawEntities{
		DrawMouse: false,
	})

	println("Running the model")
	mod.Run()
}

func createImageResource() (Grid, error) {
	f, err := logo.Open("assets/arche-logo-text.png")
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

// Grid resource
type Grid struct {
	Data   [][]bool
	Width  int
	Height int
}
