package main

import (
	"log"

	"github.com/ByteArena/box2d"
	b2d "github.com/mlange-42/arche-demo/box2d"
	"github.com/mlange-42/arche-demo/common"
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

	grav := box2d.MakeB2Vec2(0.0, 50.0)
	world := box2d.MakeB2World(grav)
	boxWorld := b2d.BoxWorld{
		World: &world,
	}
	ecs.AddResource(&game.Model.World, &boxWorld)

	ecs.AddResource(&game.Model.World, &game.Screen)
	ecs.AddResource(&game.Model.World, &game.Mouse)

	images, err := b2d.NewImages()
	if err != nil {
		println("unable to load image: ", err.Error())
		panic(err)
	}
	ecs.AddResource(&game.Model.World, &images)

	game.Model.AddSystem(&b2d.InitEntities{
		Count:       120,
		Restitution: 0.8,
	})
	game.Model.AddSystem(&b2d.Physics{
		MinFleeDistance: 50,
		MaxFleeDistance: 200,
		ForceScale:      10,
	})
	game.Model.AddSystem(&b2d.B2Physics{})

	game.Model.AddUISystem(&b2d.ManagePause{})
	game.Model.AddUISystem(&b2d.DrawEntities{})

	game.Initialize()
	if err := game.Run(); err != nil {
		log.Fatal(err)
	}
}
