package main

import (
	"embed"
	"image/png"
	"time"

	"github.com/mlange-42/arche-model/model"
	"github.com/mlange-42/arche/ecs"
)

var done chan struct{}

var cvs *Canvas
var mod *model.Model

// This specifies how long a delay between calls to 'render'.     To get Frame Rate,   1s / renderDelay
var renderDelay time.Duration = 20 * time.Millisecond

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

	cvs, _ = NewCanvas(false)
	//cvs.Create(int(js.Global().Get("innerWidth").Float()*0.9), int(js.Global().Get("innerHeight").Float()*0.9))
	cvs.Create(grid.Width, grid.Height)
	ecs.AddResource(&mod.World, cvs)

	mod.AddSystem(&InitEntities{})
	mod.AddSystem(&MoveEntities{})
	mod.AddUISystem(&DrawEntities{})

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
