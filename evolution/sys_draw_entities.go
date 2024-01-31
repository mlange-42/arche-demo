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
	grass  generic.Resource[Grass]
	canvas generic.Resource[common.EbitenImage]
	filter generic.Filter2[Position, Color]
	image  *image.RGBA
	eimage *ebiten.Image
}

// InitializeUI the system
func (s *UISysDrawEntities) InitializeUI(world *ecs.World) {
	s.grass = generic.NewResource[Grass](world)
	s.canvas = generic.NewResource[common.EbitenImage](world)
	s.filter = *generic.NewFilter2[Position, Color]()

	img := s.canvas.Get()
	s.image = image.NewRGBA(img.Image.Bounds())
	s.eimage = ebiten.NewImage(img.Image.Bounds().Dx(), img.Image.Bounds().Dy())
}

// UpdateUI the system
func (s *UISysDrawEntities) UpdateUI(world *ecs.World) {
	scale := s.grass.Get().Scale

	transp := color.RGBA{0, 0, 0, 0}

	canvas := s.canvas.Get()
	screen := canvas.Image

	// Clear the image
	draw.Draw(s.image, s.image.Bounds(), &image.Uniform{transp}, image.Point{}, draw.Src)

	// Draw pixel entities
	query := s.filter.Query(world)
	for query.Next() {
		pos, col := query.Get()

		s.image.SetRGBA(int(pos.X), int(pos.Y), col.Color)
	}
	s.eimage.WritePixels(s.image.Pix)

	geom := ebiten.GeoM{}
	geom.Scale(float64(scale), float64(scale))
	op := ebiten.DrawImageOptions{
		GeoM:   geom,
		Filter: ebiten.FilterNearest,
		Blend:  ebiten.BlendSourceOver,
	}
	screen.DrawImage(s.eimage, &op)
}

// PostUpdateUI the system
func (s *UISysDrawEntities) PostUpdateUI(world *ecs.World) {}

// FinalizeUI the system
func (s *UISysDrawEntities) FinalizeUI(world *ecs.World) {}
