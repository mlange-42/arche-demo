package ants

import (
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// SysTraceDecay is a system that decays ant traces.
type SysTraceDecay struct {
	Persistence float64

	filter generic.Filter1[Trace]
}

// Initialize the system
func (s *SysTraceDecay) Initialize(world *ecs.World) {
	s.filter = *generic.NewFilter1[Trace]()
}

// Update the system
func (s *SysTraceDecay) Update(world *ecs.World) {
	query := s.filter.Query(world)

	for query.Next() {
		trace := query.Get()
		trace.FromNest *= s.Persistence
		trace.FromResource *= s.Persistence
	}
}

// Finalize the system
func (s *SysTraceDecay) Finalize(world *ecs.World) {}
