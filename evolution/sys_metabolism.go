package evolution

import (
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// SysMetabolism is a system that reduces the energy of grazers due to metabolism.
type SysMetabolism struct {
	RateGrazing   float32
	RateSearching float32

	filterGraze  generic.Filter1[Energy]
	filterSearch generic.Filter1[Energy]
}

// Initialize the system
func (s *SysMetabolism) Initialize(world *ecs.World) {
	s.filterGraze = *generic.NewFilter1[Energy]().With(generic.T[Grazing]())
	s.filterSearch = *generic.NewFilter1[Energy]().With(generic.T[Searching]())
}

// Update the system
func (s *SysMetabolism) Update(world *ecs.World) {
	query := s.filterGraze.Query(world)
	for query.Next() {
		en := query.Get()
		en.Energy -= s.RateGrazing
	}
	query = s.filterSearch.Query(world)
	for query.Next() {
		en := query.Get()
		en.Energy -= s.RateSearching
	}
}

// Finalize the system
func (s *SysMetabolism) Finalize(world *ecs.World) {}
