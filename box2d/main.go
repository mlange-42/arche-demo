package main

import (
	"math"
	"syscall/js"

	"github.com/ByteArena/box2d"
	"github.com/mlange-42/arche-demo/common"
	"github.com/mlange-42/arche-model/model"
	"github.com/mlange-42/arche/ecs"
)

var cvs *common.Canvas
var mod *model.Model

func main() {
	mod = model.New()
	mod.FPS = 60
	mod.TPS = 60

	removeLoadingScreen("loading")

	grav := box2d.MakeB2Vec2(0.0, 50.0)
	world := box2d.MakeB2World(grav)
	boxWorld := BoxWorld{
		World: &world,
	}
	ecs.AddResource(&mod.World, &boxWorld)

	cvs, _ = common.NewCanvas("canvas-container", false)
	cvs.Create(int(math.Min(js.Global().Get("innerWidth").Float(), 880)), 480)
	ecs.AddResource(&mod.World, cvs)

	listener := MouseListener{}
	cvs.MouseListener = &listener
	ecs.AddResource(&mod.World, &listener)

	mod.AddSystem(&InitEntities{
		Count:       100,
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

func removeLoadingScreen(id string) {
	window := js.Global()
	doc := window.Get("document")
	elem := doc.Call("getElementById", id)
	elem.Call("remove")
}
