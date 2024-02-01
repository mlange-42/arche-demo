package evolution

import (
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// SysMetabolism is a system that reduces the energy of grazers due to metabolism.
type SysMetabolism struct {
	RateGrazing   float32
	RateSearching float32

	filter generic.Filter2[Energy, Activity]
}

// Initialize the system
func (s *SysMetabolism) Initialize(world *ecs.World) {
	s.filter = *generic.NewFilter2[Energy, Activity]()
}

// Update the system
func (s *SysMetabolism) Update(world *ecs.World) {
	query := s.filter.Query(world)
	for query.Next() {
		en, act := query.Get()
		if act.IsGrazing {
			en.Energy -= s.RateGrazing
		} else {
			en.Energy -= s.RateSearching
		}
	}
}

// Finalize the system
func (s *SysMetabolism) Finalize(world *ecs.World) {}
