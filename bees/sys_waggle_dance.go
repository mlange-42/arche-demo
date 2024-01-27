package bees

import (
	"github.com/mlange-42/arche-model/resource"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// SysWaggleDance is a system that handles bees doing a waggle dance at the hive ([ActWaggleDance]).
//
// Switches activity to [ActInHive] after a certain time of dancing.
// Duration of dancing depends on the resource load the bee brought back.
type SysWaggleDance struct {
	MinDanceDuration int64
	MaxDanceDuration int64

	time   generic.Resource[resource.Tick]
	filter generic.Filter1[ActWaggleDance]

	exchangeStop generic.Exchange

	toStop []ecs.Entity
}

// Initialize the system
func (s *SysWaggleDance) Initialize(world *ecs.World) {
	s.time = generic.NewResource[resource.Tick](world)
	s.filter = *generic.NewFilter1[ActWaggleDance]()

	s.exchangeStop = *generic.NewExchange(world).
		Adds(generic.T[ActInHive]()).
		Removes(generic.T[ActWaggleDance]())

	s.toStop = make([]ecs.Entity, 0, 64)
}

// Update the system
func (s *SysWaggleDance) Update(world *ecs.World) {
	tick := s.time.Get().Tick

	minDance := float64(s.MinDanceDuration)
	rangeDance := float64(s.MaxDanceDuration - s.MinDanceDuration)

	query := s.filter.Query(world)
	for query.Next() {
		wag := query.Get()

		if wag.End <= 0 {
			wag.End = tick + int64(minDance+rangeDance*wag.Load)
			continue
		}

		if tick > wag.End {
			s.toStop = append(s.toStop, query.Entity())
		}
	}

	for _, e := range s.toStop {
		s.exchangeStop.Exchange(e)
	}

	s.toStop = s.toStop[:0]
}

// Finalize the system
func (s *SysWaggleDance) Finalize(world *ecs.World) {}
