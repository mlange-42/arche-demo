package evolution

import (
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// SysMortality is a system that removes entities with energy below zero.
type SysMortality struct {
	filter generic.Filter1[Energy]

	toRemove []ecs.Entity
}

// Initialize the system
func (s *SysMortality) Initialize(world *ecs.World) {
	s.filter = *generic.NewFilter1[Energy]()

	s.toRemove = make([]ecs.Entity, 0, 16)
}

// Update the system
func (s *SysMortality) Update(world *ecs.World) {
	query := s.filter.Query(world)
	for query.Next() {
		en := query.Get()
		if en.Energy < 0 {
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
