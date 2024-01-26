package main

import (
	"math/rand"

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
	exchangeWaggle generic.Exchange
	returnMap      generic.Map1[ActReturn]
	waggleMap      generic.Map1[ActWaggleDance]
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
	s.exchangeWaggle = *generic.NewExchange(world).
		Adds(generic.T[ActWaggleDance]()).
		Removes(generic.T[ActReturn]())

	s.returnMap = generic.NewMap1[ActReturn](world)
	s.waggleMap = generic.NewMap1[ActWaggleDance](world)

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
		ret := s.returnMap.Get(e)
		if ret.Load <= 0 || rand.Float64() >= ret.Load {
			s.exchangeArrive.Exchange(e)
			continue
		}
		trg := ret.Source
		dist := distance(trg.X, trg.Y, ret.Target.X, ret.Target.Y)
		load := ret.Load
		bene := ret.Load / (dist + 1.0)

		s.exchangeWaggle.Exchange(e)
		wag := s.waggleMap.Get(e)
		wag.End = -1
		wag.Target = trg
		wag.Load = load
		wag.Benefit = bene
	}

	s.toArrive = s.toArrive[:0]
}

// Finalize the system
func (s *SysReturning) Finalize(world *ecs.World) {}
