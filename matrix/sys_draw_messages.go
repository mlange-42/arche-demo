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

	col := color.RGBA{R: 50, G: 120, B: 60, A: 255}

	query := s.filter.Query(world)
	for query.Next() {
		pos, force := query.Get()
		ln := len(messages[force.Message])
		var w float32 = 1.0
		if force.TickDone >= 0 {
			w = 2
		}
		x := float32((pos.X+1)*grid.ColumnWidth - 1)
		y := float32((pos.Y+1)*grid.LineHeight + grid.LineHeight/4)
		vector.StrokeLine(img,
			x, y,
			x+float32(ln*grid.ColumnWidth+1), y,
			w, col, false,
		)
	}
}

// PostUpdateUI the system
func (s *UISysDrawMessages) PostUpdateUI(world *ecs.World) {}

// FinalizeUI the system
func (s *UISysDrawMessages) FinalizeUI(world *ecs.World) {}
