package main

import (
	"math"
	"math/rand"

	"github.com/mlange-42/arche-model/resource"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// SysFollowing system
type SysFollowing struct {
	MaxRotation      float64
	ScoutProbability float64

	params         generic.Resource[Params]
	time           generic.Resource[resource.Tick]
	patches        generic.Resource[Patches]
	filter         generic.Filter4[Position, Direction, ActFollow, Random256]
	exchangeForage generic.Exchange
	exchangeReturn generic.Exchange
	exchangeScout  generic.Exchange
	forageMap      generic.Map1[ActForage]
	followMap      generic.Map1[ActFollow]
	posMap         generic.Map1[Position]
	homeMap        generic.Map[HomeHive]
	returnMap      generic.Map2[Position, ActReturn]
	scoutMap       generic.Map1[ActScout]

	toArrive []ecs.Entity
}

// Initialize the system
func (s *SysFollowing) Initialize(world *ecs.World) {
	s.params = generic.NewResource[Params](world)
	s.time = generic.NewResource[resource.Tick](world)
	s.patches = generic.NewResource[Patches](world)
	s.filter = *generic.NewFilter4[Position, Direction, ActFollow, Random256]()

	s.exchangeForage = *generic.NewExchange(world).
		Adds(generic.T[ActForage]()).
		Removes(generic.T[ActFollow]())
	s.exchangeReturn = *generic.NewExchange(world).
		Adds(generic.T[ActReturn]()).
		Removes(generic.T[ActFollow]())
	s.exchangeScout = *generic.NewExchange(world).
		Adds(generic.T[ActScout]()).
		Removes(generic.T[ActFollow]())

	s.forageMap = generic.NewMap1[ActForage](world)
	s.followMap = generic.NewMap1[ActFollow](world)
	s.posMap = generic.NewMap1[Position](world)
	s.homeMap = generic.NewMap[HomeHive](world)
	s.returnMap = generic.NewMap2[Position, ActReturn](world)
	s.scoutMap = generic.NewMap1[ActScout](world)

	s.toArrive = make([]ecs.Entity, 0, 32)
}

// Update the system
func (s *SysFollowing) Update(world *ecs.World) {
	tick := s.time.Get().Tick
	tickMod := tick % 4

	maxSpeed := s.params.Get().MaxBeeSpeed
	maxAng := (s.MaxRotation * math.Pi / 180.0) / 2

	patches := s.patches.Get()
	rad := float64(patches.CellSize) / 2.0

	query := s.filter.Query(world)
	for query.Next() {
		pos, dir, follow, r256 := query.Get()
		trg := follow.Target

		if tickMod == int64(r256.Value)%4 {
			dx := trg.X - pos.X
			dy := trg.Y - pos.Y

			dir.X, dir.Y, _ = norm(dx, dy)
			dir.X, dir.Y = rotate(dir.X, dir.Y, rand.Float64()*2*maxAng-maxAng)

			if dx*dx+dy*dy < rad {
				s.toArrive = append(s.toArrive, query.Entity())
				continue
			}
		}

		pos.X += dir.X * maxSpeed
		pos.Y += dir.Y * maxSpeed
	}

	for _, e := range s.toArrive {
		follow := s.followMap.Get(e)
		trg := follow.Target
		cx, cy := patches.ToCell(trg.X, trg.Y)

		patchEntity := patches.Patches[cx][cy]

		if !patchEntity.IsZero() {
			s.exchangeForage.Exchange(e)
			forage := s.forageMap.Get(e)
			forage.Start = tick
			continue
		}

		if rand.Float64() < s.ScoutProbability {
			s.exchangeScout.Exchange(e)
			scout := s.scoutMap.Get(e)
			scout.Start = tick
			continue
		}

		s.exchangeReturn.Exchange(e)
		pos, ret := s.returnMap.Get(e)
		home := s.homeMap.GetRelation(e)
		hPos := s.posMap.Get(home)
		ret.Target = *hPos
		ret.Source = *pos
		ret.Load = 0
	}

	s.toArrive = s.toArrive[:0]
}

// Finalize the system
func (s *SysFollowing) Finalize(world *ecs.World) {}
