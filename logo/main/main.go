package main

import (
	"image/png"

	"github.com/mlange-42/arche-demo/common"
	"github.com/mlange-42/arche-demo/logo"
	"github.com/mlange-42/arche-model/model"
	"github.com/mlange-42/arche/ecs"
)

var cvs *common.Canvas
var mod *model.Model

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

	cvs, _ = common.NewCanvas("canvas-container", 880, 480, true)

	image := logo.Image{Image: cvs.Image, Width: cvs.Width, Height: cvs.Height, Redraw: cvs.Redraw}
	ecs.AddResource(&mod.World, &image)

	listener := common.PauseMouseListener{}
	cvs.MouseListener = &listener
	ecs.AddResource(&mod.World, &listener)

	mod.AddSystem(&logo.InitEntities{})

	mod.AddSystem(&logo.MoveEntities{
		MaxSpeed: 10,
		MaxAcc:   0.08, MaxAccFlee: 0.1,
		MinFleeDistance: 50,
		MaxFleeDistance: 200,
		Damp:            0.975})

	mod.AddUISystem(&logo.ManagePause{})

	mod.AddUISystem(&logo.DrawEntities{})

	println("Running the model")
	mod.Run()
}

func createImageResource() (logo.Grid, error) {
	f, err := logo.Logo.Open("arche-logo-text.png")
	if err != nil {
		return logo.Grid{}, err
	}
	defer f.Close()
	img, err := png.Decode(f)
	if err != nil {
		return logo.Grid{}, err
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

	return logo.Grid{
		Data:   data,
		Width:  w,
		Height: h,
	}, nil
}
