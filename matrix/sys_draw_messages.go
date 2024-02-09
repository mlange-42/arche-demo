package matrix

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/mlange-42/arche-demo/common"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// UISysDrawLetters is a system that highlights messages.
type UISysDrawMessages struct {
	SecretKey ebiten.Key

	canvas   generic.Resource[common.EbitenImage]
	grid     generic.Resource[LetterGrid]
	messages generic.Resource[Messages]
	filter   generic.Filter2[Position, LetterForcer]
}

// InitializeUI the system
func (s *UISysDrawMessages) InitializeUI(world *ecs.World) {
	s.canvas = generic.NewResource[common.EbitenImage](world)
	s.grid = generic.NewResource[LetterGrid](world)
	s.messages = generic.NewResource[Messages](world)
	s.filter = *generic.NewFilter2[Position, LetterForcer]()
}

// UpdateUI the system
func (s *UISysDrawMessages) UpdateUI(world *ecs.World) {
	if !ebiten.IsKeyPressed(s.SecretKey) {
		return
	}

	messages := s.messages.Get().messages
	grid := s.grid.Get()
	canvas := s.canvas.Get()
	img := canvas.Image

	col := color.RGBA{R: 50, G: 127, B: 60, A: 255}

	query := s.filter.Query(world)
	for query.Next() {
		pos, force := query.Get()
		ln := len(messages[force.Message])
		vector.StrokeRect(img,
			float32((pos.X+1)*grid.ColumnWidth-2), float32(pos.Y*grid.LineHeight+grid.LineHeight/4),
			float32(ln*grid.ColumnWidth+2), float32(grid.LineHeight), 1, col, false,
		)
	}
}

// PostUpdateUI the system
func (s *UISysDrawMessages) PostUpdateUI(world *ecs.World) {}

// FinalizeUI the system
func (s *UISysDrawMessages) FinalizeUI(world *ecs.World) {}
