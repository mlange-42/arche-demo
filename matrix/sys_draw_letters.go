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
	canvas      generic.Resource[common.EbitenImage]
	grid        generic.Resource[LetterGrid]
	filterMover generic.Filter2[Position, Letter]
	filterFader generic.Filter3[Position, Letter, Fader]
}

// InitializeUI the system
func (s *UISysDrawLetters) InitializeUI(world *ecs.World) {
	s.canvas = generic.NewResource[common.EbitenImage](world)
	s.grid = generic.NewResource[LetterGrid](world)
	s.filterMover = *generic.NewFilter2[Position, Letter]().With(generic.T[Mover]())
	s.filterFader = *generic.NewFilter3[Position, Letter, Fader]()
}

// UpdateUI the system
func (s *UISysDrawLetters) UpdateUI(world *ecs.World) {
	grid := s.grid.Get()
	canvas := s.canvas.Get()
	img := canvas.Image

	col := color.RGBA{R: 100, G: 255, B: 120, A: 255}
	img.Clear()

	queryFader := s.filterFader.Query(world)
	for queryFader.Next() {
		pos, let, fad := queryFader.Get()
		v := uint8(160 * fad.Intensity)
		color := color.RGBA{R: 0, G: v, B: 0, A: 255}
		text.Draw(img, string(let.Letter), fontFaces[let.Size], (pos.X+1)*grid.ColumnWidth, pos.Y*grid.LineHeight, color)
	}

	queryMover := s.filterMover.Query(world)
	for queryMover.Next() {
		pos, let := queryMover.Get()
		text.Draw(img, string(let.Letter), fontFaces[let.Size], (pos.X+1)*grid.ColumnWidth, pos.Y*grid.LineHeight, col)
	}
}

// PostUpdateUI the system
func (s *UISysDrawLetters) PostUpdateUI(world *ecs.World) {}

// FinalizeUI the system
func (s *UISysDrawLetters) FinalizeUI(world *ecs.World) {}
