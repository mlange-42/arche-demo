package box2d

import (
	"log"

	"github.com/ByteArena/box2d"
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

	grav := box2d.MakeB2Vec2(0.0, 50.0)
	world := box2d.MakeB2World(grav)
	boxWorld := BoxWorld{
		World: &world,
	}
	ecs.AddResource(&game.Model.World, &boxWorld)

	ecs.AddResource(&game.Model.World, &game.Screen)
	ecs.AddResource(&game.Model.World, &game.Mouse)

	images, err := NewImages()
	if err != nil {
		println("unable to load image: ", err.Error())
		panic(err)
	}
	ecs.AddResource(&game.Model.World, &images)

	game.Model.AddSystem(&SysInitEntities{
		Count:       120,
		Restitution: 0.8,
	})
	game.Model.AddSystem(&SysPhysics{
		MinFleeDistance: 50,
		MaxFleeDistance: 200,
		ForceScale:      10,
	})
	game.Model.AddSystem(&SysB2Physics{})

	game.Model.AddUISystem(&UISysManagePause{})
	game.Model.AddUISystem(&UISysDrawEntities{})

	game.Initialize()
	if err := game.Run(); err != nil {
		log.Fatal(err)
	}
}
