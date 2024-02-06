package boids

import (
	"embed"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

//go:embed boid.png
var images embed.FS

// Images resource for rendering.
type Images struct {
	Boid *ebiten.Image
}

func NewImages() Images {
	boids, _, err := ebitenutil.NewImageFromFileSystem(images, "boid.png")
	if err != nil {
		panic("can't load image boid.png")
	}
	return Images{
		Boid: boids,
	}
}
