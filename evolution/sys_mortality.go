package evolution

import (
	"github.com/mlange-42/arche-model/resource"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// SysMortality is a system that removes entities with energy below zero.
type SysMortality struct {
	MaxAge int64

	time   generic.Resource[resource.Tick]
	filter generic.Filter2[Energy, Age]

	toRemove []ecs.Entity
}

// Initialize the system
func (s *SysMortality) Initialize(world *ecs.World) {
	s.time = generic.NewResource[resource.Tick](world)
	s.filter = *generic.NewFilter2[Energy, Age]()

	s.toRemove = make([]ecs.Entity, 0, 16)
}

// Update the system
func (s *SysMortality) Update(world *ecs.World) {
	tick := s.time.Get().Tick

	query := s.filter.Query(world)
	for query.Next() {
		en, age := query.Get()
		if en.Energy < 0 || tick > age.TickOfBirth+s.MaxAge {
			s.toRemove = append(s.toRemove, query.Entity())
		}
	}

	for _, e := range s.toRemove {
		world.RemoveEntity(e)
	}

	s.toRemove = s.toRemove[:0]
}

// Finalize the system
func (s *SysMortality) Finalize(world *ecs.World) {}
