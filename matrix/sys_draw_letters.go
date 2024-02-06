package matrix

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/mlange-42/arche-demo/common"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// UISysDrawLetters is a system that draws boids.
type UISysDrawLetters struct {
	canvas generic.Resource[common.EbitenImage]
	filter generic.Filter2[Position, Letter]
}

// InitializeUI the system
func (s *UISysDrawLetters) InitializeUI(world *ecs.World) {
	s.canvas = generic.NewResource[common.EbitenImage](world)
	s.filter = *generic.NewFilter2[Position, Letter]()
}

// UpdateUI the system
func (s *UISysDrawLetters) UpdateUI(world *ecs.World) {
	//images := s.images.Get()
	canvas := s.canvas.Get()
	img := canvas.Image

	col := color.RGBA{R: 0, G: 255, B: 0, A: 255}

	img.Clear()

	query := s.filter.Query(world)
	for query.Next() {
		pos, let := query.Get()
		text.Draw(img, string(let.Letter), fontFaces[let.Size], int(pos.X), int(pos.Y), col)
	}
}

// PostUpdateUI the system
func (s *UISysDrawLetters) PostUpdateUI(world *ecs.World) {}

// FinalizeUI the system
func (s *UISysDrawLetters) FinalizeUI(world *ecs.World) {}
