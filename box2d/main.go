package main

import (
	"embed"
	"image"
	"image/draw"
	"image/png"
	"math"
	"syscall/js"

	"github.com/ByteArena/box2d"
	"github.com/mlange-42/arche-demo/common"
	"github.com/mlange-42/arche-model/model"
	"github.com/mlange-42/arche/ecs"
)

var cvs *common.Canvas
var mod *model.Model

//go:embed circle.png
var circle embed.FS

func main() {
	mod = model.New()
	mod.FPS = 30
	mod.TPS = 60

	common.RemoveElementByID("loading")

	grav := box2d.MakeB2Vec2(0.0, 50.0)
	world := box2d.MakeB2World(grav)
	boxWorld := BoxWorld{
		World: &world,
	}
	ecs.AddResource(&mod.World, &boxWorld)

	cvs, _ = common.NewCanvas("canvas-container", false)
	cvs.Create(int(math.Min(js.Global().Get("innerWidth").Float(), 880)), 480)

	image := Image{Image: cvs.Image, Width: cvs.Width, Height: cvs.Height, Redraw: cvs.Redraw}
	ecs.AddResource(&mod.World, &image)

	listener := MouseListener{}
	cvs.MouseListener = &listener
	ecs.AddResource(&mod.World, &listener)

	images, err := createImagesResource()
	if err != nil {
		println("unable to load image: ", err.Error())
		panic(err)
	}
	ecs.AddResource(&mod.World, &images)

	mod.AddSystem(&InitEntities{
		Count:       80,
		Restitution: 0.8,
	})
	mod.AddSystem(&Physics{
		MinFleeDistance: 50,
		MaxFleeDistance: 200,
		ForceScale:      10,
	})
	mod.AddSystem(&Box2DPhysics{})

	mod.AddUISystem(&ManagePause{})
	mod.AddUISystem(&DrawEntities{})

	println("Running the model")
	mod.Run()
}

func createImagesResource() (Images, error) {
	f, err := circle.Open("circle.png")
	if err != nil {
		return Images{}, err
	}
	defer f.Close()
	src, err := png.Decode(f)
	if err != nil {
		return Images{}, err
	}

	b := src.Bounds()
	img := image.NewRGBA(image.Rect(0, 0, b.Dx(), b.Dy()))
	draw.Draw(img, img.Bounds(), src, b.Min, draw.Src)

	return Images{
		Circle: img,
	}, nil
}
