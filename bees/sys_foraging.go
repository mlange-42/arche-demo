package main

import (
	"github.com/mlange-42/arche-model/resource"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// SysForaging system
type SysForaging struct {
	MaxForagingTime int64
	patches         generic.Resource[Patches]
	time            generic.Resource[resource.Tick]
	filter          generic.Filter1[ActForage]

	exchangeReturn generic.Exchange

	toReturn []ecs.Entity
}

// Initialize the system
func (s *SysForaging) Initialize(world *ecs.World) {
	s.patches = generic.NewResource[Patches](world)
	s.time = generic.NewResource[resource.Tick](world)

	s.filter = *generic.NewFilter1[ActForage]()

	s.exchangeReturn = *generic.NewExchange(world).
		Removes(generic.T[ActForage]())

	s.toReturn = make([]ecs.Entity, 0, 64)
}

// Update the system
func (s *SysForaging) Update(world *ecs.World) {
	tick := s.time.Get().Tick

	query := s.filter.Query(world)
	for query.Next() {
		forage := query.Get()

		if tick > forage.Start+s.MaxForagingTime {
			s.toReturn = append(s.toReturn, query.Entity())
			continue
		}
	}

	for _, e := range s.toReturn {
		s.exchangeReturn.Exchange(e)
	}

	s.toReturn = s.toReturn[:0]
}

// Finalize the system
func (s *SysForaging) Finalize(world *ecs.World) {}
