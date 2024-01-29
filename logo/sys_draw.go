package logo

import (
	"image"
	"image/color"
	"image/draw"

	"github.com/mlange-42/arche-demo/common"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// DrawEntities is a system that draws entities as white pixels on an [Image] resource.
type DrawEntities struct {
	canvas generic.Resource[common.EbitenImage]
	filter generic.Filter1[Position]
	image  *image.RGBA
}

// InitializeUI the system
func (s *DrawEntities) InitializeUI(world *ecs.World) {
	s.canvas = generic.NewResource[common.EbitenImage](world)
	s.filter = *generic.NewFilter1[Position]()

	img := s.canvas.Get()
	s.image = image.NewRGBA(img.Image.Bounds())
}

// UpdateUI the system
func (s *DrawEntities) UpdateUI(world *ecs.World) {
	black := color.RGBA{0, 0, 0, 255}
	white := color.RGBA{255, 255, 255, 255}

	canvas := s.canvas.Get()
	img := canvas.Image

	// Clear the image
	draw.Draw(s.image, s.image.Bounds(), &image.Uniform{black}, image.Point{}, draw.Src)

	// Draw pixel entities
	query := s.filter.Query(world)
	for query.Next() {
		pos := query.Get()

		s.image.SetRGBA(int(pos.X), int(pos.Y), white)
	}
	img.WritePixels(s.image.Pix)
}

// PostUpdateUI the system
func (s *DrawEntities) PostUpdateUI(world *ecs.World) {}

// FinalizeUI the system
func (s *DrawEntities) FinalizeUI(world *ecs.World) {}
