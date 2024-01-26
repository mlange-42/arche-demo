package main

import (
	"math/rand"

	"github.com/mlange-42/arche-model/resource"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// SysHiveDecisions system
type SysHiveDecisions struct {
	ReleaseInterval  int64
	ReleaseCount     int
	ScoutProbability float64

	DanceSamples int

	time         generic.Resource[resource.Tick]
	patches      generic.Resource[Patches]
	hiveFilter   generic.Filter1[Random256]
	idleFilter   generic.Filter0
	waggleFilter generic.Filter1[ActWaggleDance]

	exchangeScout  generic.Exchange
	scoutMap       generic.Map1[ActScout]
	exchangeFollow generic.Exchange
	followMap      generic.Map1[ActFollow]

	hives   []ecs.Entity
	dances  []waggleInfo
	toLeave []ecs.Entity
}

type waggleInfo struct {
	Target  Position
	Benefit float64
}

// Initialize the system
func (s *SysHiveDecisions) Initialize(world *ecs.World) {
	s.time = generic.NewResource[resource.Tick](world)
	s.hiveFilter = *generic.NewFilter1[Random256]().With(generic.T[Hive]())
	s.idleFilter = *generic.NewFilter0().With(generic.T2[ActInHive, HomeHive]()...).WithRelation(generic.T[HomeHive]())
	s.waggleFilter = *generic.NewFilter1[ActWaggleDance]().With(generic.T[HomeHive]()).WithRelation(generic.T[HomeHive]())

	s.exchangeScout = *generic.NewExchange(world).Adds(generic.T[ActScout]()).Removes(generic.T[ActInHive]())
	s.scoutMap = generic.NewMap1[ActScout](world)

	s.exchangeFollow = *generic.NewExchange(world).Adds(generic.T[ActFollow]()).Removes(generic.T[ActInHive]())
	s.followMap = generic.NewMap1[ActFollow](world)

	s.hives = make([]ecs.Entity, 0, 16)
	s.dances = make([]waggleInfo, 0, 16)
	s.toLeave = make([]ecs.Entity, 0, 64)
}

// Update the system
func (s *SysHiveDecisions) Update(world *ecs.World) {
	tick := s.time.Get().Tick
	tickMod := tick % s.ReleaseInterval

	query := s.hiveFilter.Query(world)
	for query.Next() {
		r256 := query.Get()
		if tickMod == int64(r256.Value)%s.ReleaseInterval {
			s.hives = append(s.hives, query.Entity())
		}
	}

	for _, hiveEntity := range s.hives {
		waggleQuery := s.waggleFilter.Query(world, hiveEntity)
		for waggleQuery.Next() {
			wag := waggleQuery.Get()
			s.dances = append(s.dances, waggleInfo{Target: wag.Target, Benefit: wag.Benefit})
		}

		idleQuery := s.idleFilter.Query(world, hiveEntity)
		totalCnt := idleQuery.Count()
		cnt := 0
		for idleQuery.Next() {
			if cnt >= s.ReleaseCount {
				break
			}
			s.toLeave = append(s.toLeave, idleQuery.Entity())
			cnt++
		}
		if cnt < totalCnt {
			idleQuery.Close()
		}

		for _, e := range s.toLeave {
			if len(s.dances) < 2 || rand.Float64() < s.ScoutProbability {
				s.exchangeScout.Exchange(e)
				scout := s.scoutMap.Get(e)
				scout.Start = tick
				continue
			}
			wag, ok := s.selectDance()
			if !ok {
				panic("unable to select a waggle dance")
			}
			s.exchangeFollow.Exchange(e)
			follow := s.followMap.Get(e)
			follow.Target = wag.Target
		}

		s.dances = s.dances[:0]
		s.toLeave = s.toLeave[:0]
	}

	s.hives = s.hives[:0]
}

func (s *SysHiveDecisions) selectDance() (waggleInfo, bool) {
	var best waggleInfo
	maxBenefit := -1.0

	ln := len(s.dances)
	for i := 0; i < s.DanceSamples; i++ {
		dance := s.dances[rand.Intn(ln)]
		if dance.Benefit > maxBenefit {
			best = dance
			maxBenefit = dance.Benefit
		}
	}

	return best, maxBenefit >= 0
}

// Finalize the system
func (s *SysHiveDecisions) Finalize(world *ecs.World) {}
