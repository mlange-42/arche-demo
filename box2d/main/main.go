package main

import (
	"github.com/ByteArena/box2d"
	b2d "github.com/mlange-42/arche-demo/box2d"
	"github.com/mlange-42/arche-demo/common"
	"github.com/mlange-42/arche-model/model"
	"github.com/mlange-42/arche/ecs"
)

var cvs *common.Canvas
var mod *model.Model

func main() {
	mod = model.New()
	mod.FPS = 30
	mod.TPS = 60

	grav := box2d.MakeB2Vec2(0.0, 50.0)
	world := box2d.MakeB2World(grav)
	boxWorld := b2d.BoxWorld{
		World: &world,
	}
	ecs.AddResource(&mod.World, &boxWorld)

	cvs, _ = common.NewCanvas("canvas-container", 880, 480, true)

	image := b2d.Image{Image: cvs.Image, Width: cvs.Width, Height: cvs.Height, Redraw: cvs.Redraw}
	ecs.AddResource(&mod.World, &image)

	listener := common.PauseMouseListener{}
	cvs.MouseListener = &listener
	ecs.AddResource(&mod.World, &listener)

	images, err := b2d.NewImages()
	if err != nil {
		println("unable to load image: ", err.Error())
		panic(err)
	}
	ecs.AddResource(&mod.World, &images)

	mod.AddSystem(&b2d.InitEntities{
		Count:       80,
		Restitution: 0.8,
	})
	mod.AddSystem(&b2d.Physics{
		MinFleeDistance: 50,
		MaxFleeDistance: 200,
		ForceScale:      10,
	})
	mod.AddSystem(&b2d.B2Physics{})

	mod.AddUISystem(&b2d.ManagePause{})
	mod.AddUISystem(&b2d.DrawEntities{})

	println("Running the model")
	mod.Run()
}
