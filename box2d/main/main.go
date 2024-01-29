package main

import (
	"log"

	"github.com/ByteArena/box2d"
	"github.com/hajimehoshi/ebiten/v2"
	b2d "github.com/mlange-42/arche-demo/box2d"
	"github.com/mlange-42/arche-model/model"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

const (
	screenWidth  = 880
	screenHeight = 480
)

// Game container
type Game struct {
	mod      *model.Model
	boxWorld box2d.B2World
	images   b2d.Images
	image    generic.Resource[b2d.Image]
}

// NewGame returns a new game
func NewGame() Game {
	game := Game{}

	game.mod = model.New()
	game.mod.FPS = 9999
	game.mod.TPS = 9999

	grav := box2d.MakeB2Vec2(0.0, 50.0)
	game.boxWorld = box2d.MakeB2World(grav)
	boxWorld := b2d.BoxWorld{
		World: &game.boxWorld,
	}
	ecs.AddResource(&game.mod.World, &boxWorld)

	image := b2d.Image{Image: ebiten.NewImage(screenWidth, screenHeight), Width: screenWidth, Height: screenHeight}
	ecs.AddResource(&game.mod.World, &image)
	game.image = generic.NewResource[b2d.Image](&game.mod.World)

	var err error
	game.images, err = b2d.NewImages()
	if err != nil {
		println("unable to load image: ", err.Error())
		panic(err)
	}
	ecs.AddResource(&game.mod.World, &game.images)

	game.mod.AddSystem(&b2d.InitEntities{
		Count:       120,
		Restitution: 0.8,
	})
	game.mod.AddSystem(&b2d.Physics{
		MinFleeDistance: 50,
		MaxFleeDistance: 200,
		ForceScale:      10,
	})
	game.mod.AddSystem(&b2d.B2Physics{})

	game.mod.AddUISystem(&b2d.ManagePause{})
	game.mod.AddUISystem(&b2d.DrawEntities{})

	return game
}

// Layout the game
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	//return int(math.Min(float64(outsideWidth), screenWidth)), int(math.Min(float64(outsideHeight), screenHeight))
	return screenWidth, screenHeight
}

// Update the game.
func (g *Game) Update() error {
	g.mod.UpdateSystems()
	return nil
}

// Draw the game.
func (g *Game) Draw(screen *ebiten.Image) {
	image := g.image.Get()
	g.mod.UpdateUISystems()

	options := ebiten.DrawImageOptions{}

	screen.DrawImage(image.Image, &options)
}

func main() {
	game := NewGame()
	game.mod.Initialize()

	ebiten.SetFullscreen(false)
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Box2D")
	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
