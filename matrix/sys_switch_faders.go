package matrix

import (
	"math/rand"

	"github.com/mlange-42/arche-model/resource"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// SysSwitchFaders is a system that occasionally changes the character of faders.
type SysSwitchFaders struct {
	MinChangeInterval int
	MaxChangeInterval int

	time    generic.Resource[resource.Tick]
	letters generic.Resource[Letters]
	filter  generic.Filter3[Letter, Fader, ForcedLetter]
}

// Initialize the system
func (s *SysSwitchFaders) Initialize(world *ecs.World) {
	s.time = generic.NewResource[resource.Tick](world)
	s.letters = generic.NewResource[Letters](world)
	s.filter = *generic.NewFilter3[Letter, Fader, ForcedLetter]()
}

// Update the system
func (s *SysSwitchFaders) Update(world *ecs.World) {
	tick := s.time.Get().Tick
	letters := s.letters.Get().Letters

	query := s.filter.Query(world)
	for query.Next() {
		let, fad, forced := query.Get()
		if forced.Active || fad.Intensity < 0 || tick < fad.NextChange {
			continue
		}
		if fad.NextChange > 0 {
			let.Letter = letters[rand.Intn(len(letters))]
		}
		fad.NextChange = tick + rand.Int63n(int64(s.MaxChangeInterval)-int64(s.MinChangeInterval)) + int64(s.MinChangeInterval)
	}
}

// Finalize the system
func (s *SysSwitchFaders) Finalize(world *ecs.World) {}
