package ants

import (
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// SysDecay is a system that decays ant traces.
type SysDecay struct {
	Persistence float64

	filter generic.Filter1[Trace]
}

// Initialize the system
func (s *SysDecay) Initialize(world *ecs.World) {
	s.filter = *generic.NewFilter1[Trace]()
}

// Update the system
func (s *SysDecay) Update(world *ecs.World) {
	query := s.filter.Query(world)

	for query.Next() {
		trace := query.Get()
		trace.FromNest *= s.Persistence
		trace.FromResource *= s.Persistence
	}
}

// Finalize the system
func (s *SysDecay) Finalize(world *ecs.World) {}
