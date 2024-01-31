package common

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mlange-42/arche-model/model"
	"github.com/mlange-42/arche/generic"
)

// Game container
type Game struct {
	Model  *model.Model
	Screen EbitenImage
	Mouse  Mouse
	width  int
	height int
	canvas *canvasHelper

	frame uint64
	speed generic.Resource[SimulationSpeed]
}

// NewGame returns a new game
func NewGame(mod *model.Model, width, height int) Game {
	return Game{
		Model:  mod,
		Screen: EbitenImage{Image: ebiten.NewImage(width, height), Width: width, Height: height},
		width:  width,
		height: height,
		canvas: newCanvasHelper(width, height),
		frame:  0,
		speed:  generic.NewResource[SimulationSpeed](&mod.World),
	}
}

// Initialize the game.
func (g *Game) Initialize() {
	ebiten.SetWindowSize(g.width, g.height)
	g.Model.Initialize()
}

// Run the game.
func (g *Game) Run() error {
	if err := ebiten.RunGame(g); err != nil {
		return err
	}
	return nil
}

// Layout the game.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.width, g.height
}

// Update the game.
func (g *Game) Update() error {
	speed := 1.0
	if g.speed.Has() {
		exp := g.speed.Get().Exponent
		speed = math.Pow(2, float64(exp))
	}
	if speed >= 1 {
		iters := int(speed)
		for i := 0; i < iters; i++ {
			g.updateMouse()
			g.Model.Update()
		}
	} else {
		iters := int(1.0 / speed)
		if g.frame%uint64(iters) == 0 {
			g.updateMouse()
			g.Model.Update()
		}
	}

	g.frame++
	return nil
}

// Draw the game.
func (g *Game) Draw(screen *ebiten.Image) {
	g.updateMouse()

	g.Screen.Image.Clear()
	g.Model.UpdateUI()

	options := ebiten.DrawImageOptions{}
	screen.DrawImage(g.Screen.Image, &options)
}

func (g *Game) updateMouse() {
	x, y := ebiten.CursorPosition()
	g.Mouse.IsInside = g.IsMouseInside()
	g.Mouse.X = x
	g.Mouse.Y = y
}

// IsMouseInside returns whether the mouse is inside the game canvas.
func (g *Game) IsMouseInside() bool {
	return g.canvas.isMouseInside()
}
