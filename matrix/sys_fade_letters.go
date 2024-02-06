package matrix

import (
	"github.com/mlange-42/arche-model/resource"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// SysMoveLetters is a system to initialize entities.
type SysFadeLetters struct {
	FadeDuration int

	time   generic.Resource[resource.Tick]
	filter generic.Filter1[Fader]

	toRemove []ecs.Entity
}

// Initialize the system
func (s *SysFadeLetters) Initialize(world *ecs.World) {
	s.time = generic.NewResource[resource.Tick](world)
	s.filter = *generic.NewFilter1[Fader]()
}

// Update the system
func (s *SysFadeLetters) Update(world *ecs.World) {
	tick := s.time.Get().Tick

	query := s.filter.Query(world)
	for query.Next() {
		fad := query.Get()
		fad.Intensity = 1.0 - float64(tick-fad.Start)/float64(s.FadeDuration)
		if fad.Intensity <= 0 {
			s.toRemove = append(s.toRemove, query.Entity())
		}
	}

	for _, e := range s.toRemove {
		world.RemoveEntity(e)
	}

	s.toRemove = s.toRemove[:0]
}

// Finalize the system
func (s *SysFadeLetters) Finalize(world *ecs.World) {}
