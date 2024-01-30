package evolution

import (
	"image"
	"image/color"
	"image/draw"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mlange-42/arche-demo/common"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// UISysDrawEntities is a system that draws entities as white pixels on an [Image] resource.
type UISysDrawEntities struct {
	canvas generic.Resource[common.EbitenImage]
	filter generic.Filter1[Position]
	image  *image.RGBA
	eimage *ebiten.Image
}

// InitializeUI the system
func (s *UISysDrawEntities) InitializeUI(world *ecs.World) {
	s.canvas = generic.NewResource[common.EbitenImage](world)
	s.filter = *generic.NewFilter1[Position]()

	img := s.canvas.Get()
	s.image = image.NewRGBA(img.Image.Bounds())
	s.eimage = ebiten.NewImage(img.Image.Bounds().Dx(), img.Image.Bounds().Dy())
}

// UpdateUI the system
func (s *UISysDrawEntities) UpdateUI(world *ecs.World) {
	transp := color.RGBA{0, 0, 0, 0}
	white := color.RGBA{255, 255, 255, 255}

	canvas := s.canvas.Get()
	screen := canvas.Image

	// Clear the image
	draw.Draw(s.image, s.image.Bounds(), &image.Uniform{transp}, image.Point{}, draw.Src)

	// Draw pixel entities
	query := s.filter.Query(world)
	for query.Next() {
		pos := query.Get()

		s.image.SetRGBA(int(pos.X), int(pos.Y), white)
	}
	s.eimage.WritePixels(s.image.Pix)

	op := ebiten.DrawImageOptions{
		Filter: ebiten.FilterNearest,
		Blend:  ebiten.BlendSourceOver,
	}
	screen.DrawImage(s.eimage, &op)
}

// PostUpdateUI the system
func (s *UISysDrawEntities) PostUpdateUI(world *ecs.World) {}

// FinalizeUI the system
func (s *UISysDrawEntities) FinalizeUI(world *ecs.World) {}