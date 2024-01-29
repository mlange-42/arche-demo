package common

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mlange-42/arche-model/model"
)

// Game container
type Game struct {
	Model  *model.Model
	Screen EbitenImage
	Mouse  Mouse
	width  int
	height int
	canvas *canvasHelper
}

// NewGame returns a new game
func NewGame(mod *model.Model, width, height int) Game {
	return Game{
		Model:  mod,
		Screen: EbitenImage{Image: ebiten.NewImage(width, height), Width: width, Height: height},
		width:  width,
		height: height,
		canvas: newCanvasHelper(),
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
	g.updateMouse()

	g.Model.UpdateSystems()
	return nil
}

// Draw the game.
func (g *Game) Draw(screen *ebiten.Image) {
	g.updateMouse()

	g.Screen.Image.Clear()
	g.Model.UpdateUISystems()

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
