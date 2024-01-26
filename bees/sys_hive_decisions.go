package main

import (
	"math/rand"

	"github.com/mlange-42/arche-model/resource"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// SysHiveDecisions system
type SysHiveDecisions struct {
	time       generic.Resource[resource.Tick]
	hiveFilter generic.Filter1[Random256]
	beesFilter generic.Filter0

	exchangeScout generic.Exchange
	scoutMap      generic.Map1[ActScout]

	toScout []ecs.Entity
}

// Initialize the system
func (s *SysHiveDecisions) Initialize(world *ecs.World) {
	s.time = generic.NewResource[resource.Tick](world)
	s.hiveFilter = *generic.NewFilter1[Random256]().With(generic.T[Hive]())
	s.beesFilter = *generic.NewFilter0().With(generic.T2[ActInHive, HomeHive]()...).WithRelation(generic.T[HomeHive]())

	s.exchangeScout = *generic.NewExchange(world).Adds(generic.T[ActScout]()).Removes(generic.T[ActInHive]())
	s.scoutMap = generic.NewMap1[ActScout](world)

	s.toScout = make([]ecs.Entity, 0, 64)
}

// Update the system
func (s *SysHiveDecisions) Update(world *ecs.World) {
	tick := s.time.Get().Tick
	tickMod := tick % 4

	query := s.hiveFilter.Query(world)
	for query.Next() {
		r256 := query.Get()

		if tickMod != int64(r256.Value)%4 {
			continue
		}

		hiveEntity := query.Entity()
		beesQuery := s.beesFilter.Query(world, hiveEntity)
		for beesQuery.Next() {
			if rand.Float64() < 0.01 {
				s.toScout = append(s.toScout, beesQuery.Entity())
			}
		}
	}

	for _, e := range s.toScout {
		s.exchangeScout.Exchange(e)
		scout := s.scoutMap.Get(e)
		scout.Start = tick
	}

	s.toScout = s.toScout[:0]
}

// Finalize the system
func (s *SysHiveDecisions) Finalize(world *ecs.World) {}
