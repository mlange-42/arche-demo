package matrix

import (
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// SysMoveLetters is a system to fade out letters over time.
type SysFadeLetters struct {
	FadeDuration int

	filter generic.Filter1[Fader]
}

// Initialize the system
func (s *SysFadeLetters) Initialize(world *ecs.World) {
	s.filter = *generic.NewFilter1[Fader]()
}

// Update the system
func (s *SysFadeLetters) Update(world *ecs.World) {
	v := 1.0 / float64(s.FadeDuration)

	query := s.filter.Query(world)
	for query.Next() {
		fad := query.Get()
		fad.Intensity -= v
	}
}

// Finalize the system
func (s *SysFadeLetters) Finalize(world *ecs.World) {}
