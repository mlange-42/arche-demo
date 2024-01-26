package main

import (
	"github.com/mlange-42/arche-model/resource"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// SysReturning system
type SysReturning struct {
	params         generic.Resource[Params]
	time           generic.Resource[resource.Tick]
	filter         generic.Filter4[Position, Direction, ActReturn, Random256]
	exchangeArrive generic.Exchange
	toArrive       []ecs.Entity
}

// Initialize the system
func (s *SysReturning) Initialize(world *ecs.World) {
	s.params = generic.NewResource[Params](world)
	s.time = generic.NewResource[resource.Tick](world)
	s.filter = *generic.NewFilter4[Position, Direction, ActReturn, Random256]().With(generic.T[ActReturn]())

	s.exchangeArrive = *generic.NewExchange(world).
		Adds(generic.T[ActInHive]()).
		Removes(generic.T[ActReturn]())

	s.toArrive = make([]ecs.Entity, 0, 64)
}

// Update the system
func (s *SysReturning) Update(world *ecs.World) {
	tick := s.time.Get().Tick
	maxSpeed := s.params.Get().MaxBeeSpeed

	query := s.filter.Query(world)
	for query.Next() {
		pos, dir, ret, r256 := query.Get()

		if tick%4 == int64(r256.Value)%4 {
			dx := ret.Target.X - pos.X
			dy := ret.Target.Y - pos.Y

			dir.X, dir.Y, _ = norm(dx, dy)

			if dx*dx+dy*dy < 4 {
				s.toArrive = append(s.toArrive, query.Entity())
				continue
			}
		}

		pos.X += dir.X * maxSpeed
		pos.Y += dir.Y * maxSpeed
	}

	for _, e := range s.toArrive {
		s.exchangeArrive.Exchange(e)
	}

	s.toArrive = s.toArrive[:0]
}

// Finalize the system
func (s *SysReturning) Finalize(world *ecs.World) {}
