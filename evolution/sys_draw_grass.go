package evolution

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mlange-42/arche-demo/common"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// UISysDrawGrass is a system that draws the [Grass] resource.
type UISysDrawGrass struct {
	canvas generic.Resource[common.EbitenImage]
	grass  generic.Resource[Grass]
	image  *image.RGBA
	eimage *ebiten.Image
}

// InitializeUI the system
func (s *UISysDrawGrass) InitializeUI(world *ecs.World) {
	s.canvas = generic.NewResource[common.EbitenImage](world)
	s.grass = generic.NewResource[Grass](world)

	grass := &s.grass.Get().Grass
	s.image = image.NewRGBA(image.Rect(0, 0, grass.Width(), grass.Height()))
	s.eimage = ebiten.NewImage(grass.Width(), grass.Height())
}

// UpdateUI the system
func (s *UISysDrawGrass) UpdateUI(world *ecs.World) {
	grassRes := s.grass.Get()
	grass := &grassRes.Grass
	scale := grassRes.Scale

	canvas := s.canvas.Get()
	screen := canvas.Image

	w, h := grass.Width(), grass.Height()

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			v := grass.Get(x, y)
			col := color.RGBA{R: 0, G: uint8(v * 80), B: 0, A: 255}
			s.image.Set(x, y, col)
		}
	}
	s.eimage.WritePixels(s.image.Pix)

	geom := ebiten.GeoM{}
	geom.Scale(grass.Cellsize()*float64(scale), grass.Cellsize()*float64(scale))
	op := ebiten.DrawImageOptions{
		GeoM:   geom,
		Filter: ebiten.FilterNearest,
	}
	screen.DrawImage(s.eimage, &op)
}

// PostUpdateUI the system
func (s *UISysDrawGrass) PostUpdateUI(world *ecs.World) {}

// FinalizeUI the system
func (s *UISysDrawGrass) FinalizeUI(world *ecs.World) {}
