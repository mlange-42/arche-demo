package matrix

import (
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// SysMoveLetters is a system to fade out letters over time.
type SysFadeLetters struct {
	FadeDuration        int
	MessageFadeDuration int

	grid   generic.Resource[LetterGrid]
	filter generic.Filter3[Position, Fader, ForcedLetter]
}

// Initialize the system
func (s *SysFadeLetters) Initialize(world *ecs.World) {
	s.grid = generic.NewResource[LetterGrid](world)
	s.filter = *generic.NewFilter3[Position, Fader, ForcedLetter]()
}

// Update the system
func (s *SysFadeLetters) Update(world *ecs.World) {
	height := s.grid.Get().Faders.Height()
	v := 1.0 / float64(s.FadeDuration)
	vm := 1.0 / float64(s.MessageFadeDuration)

	query := s.filter.Query(world)
	for query.Next() {
		pos, fad, forced := query.Get()
		if forced.Active && forced.Traversed {
			sc := float64(pos.X) / float64(height)
			fad.Intensity -= (1-sc)*vm + sc*v
		} else {
			fad.Intensity -= v
		}
	}
}

// Finalize the system
func (s *SysFadeLetters) Finalize(world *ecs.World) {}
