package matrix

import (
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// SysMoveLetters is a system to fade out letters over time.
type SysFadeLetters struct {
	FadeDuration        int
	MessageFadeDuration int

	filter generic.Filter2[Fader, ForcedLetter]
}

// Initialize the system
func (s *SysFadeLetters) Initialize(world *ecs.World) {
	s.filter = *generic.NewFilter2[Fader, ForcedLetter]()
}

// Update the system
func (s *SysFadeLetters) Update(world *ecs.World) {
	v := 1.0 / float64(s.FadeDuration)
	vm := 1.0 / float64(s.MessageFadeDuration)

	query := s.filter.Query(world)
	for query.Next() {
		fad, forced := query.Get()
		if forced.Active && forced.Traversed {
			fad.Intensity -= vm
		} else {
			fad.Intensity -= v
		}
	}
}

// Finalize the system
func (s *SysFadeLetters) Finalize(world *ecs.World) {}
