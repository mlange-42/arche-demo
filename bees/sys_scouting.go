package main

import (
	"math"
	"math/rand"

	"github.com/mlange-42/arche-model/resource"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// SysScouting system
type SysScouting struct {
	MaxRotation  float64
	MaxScoutTime int64
	canvas       generic.Resource[Image]
	patches      generic.Resource[Patches]
	params       generic.Resource[Params]
	time         generic.Resource[resource.Tick]
	filter       generic.Filter4[Position, Direction, ActScout, Random256]

	exchangeForage generic.Exchange
	exchangeReturn generic.Exchange
	forageMap      generic.Map1[ActForage]

	toForage []ecs.Entity
	toReturn []ecs.Entity
}

// Initialize the system
func (s *SysScouting) Initialize(world *ecs.World) {
	s.canvas = generic.NewResource[Image](world)
	s.patches = generic.NewResource[Patches](world)
	s.params = generic.NewResource[Params](world)
	s.time = generic.NewResource[resource.Tick](world)

	s.filter = *generic.NewFilter4[Position, Direction, ActScout, Random256]()

	s.exchangeForage = *generic.NewExchange(world).
		Adds(generic.T[ActForage]()).
		Removes(generic.T[ActScout]())
	s.exchangeReturn = *generic.NewExchange(world).
		Removes(generic.T[ActScout]())

	s.forageMap = generic.NewMap1[ActForage](world)

	s.toForage = make([]ecs.Entity, 0, 64)
	s.toReturn = make([]ecs.Entity, 0, 64)
}

// Update the system
func (s *SysScouting) Update(world *ecs.World) {
	canvas := s.canvas.Get()
	patches := s.patches.Get()

	w := float64(canvas.Width)
	h := float64(canvas.Height)

	maxSpeed := s.params.Get().MaxBeeSpeed
	maxAng := (s.MaxRotation * math.Pi / 180.0) / 2
	tick := s.time.Get().Tick

	query := s.filter.Query(world)
	for query.Next() {
		pos, dir, scout, r256 := query.Get()

		// Drawing random numbers is costly, so we do it only every 4 ticks.
		// We also check to end scouting here, as this is not necessary every tick.
		if tick%4 == int64(r256.Value)%4 {
			dir.X, dir.Y = rotate(dir.X, dir.Y, rand.Float64()*2*maxAng-maxAng)

			cx, cy := patches.ToCell(pos.X, pos.Y)
			patch := patches.Patches[cx][cy]
			if !patch.IsZero() {
				s.toForage = append(s.toForage, query.Entity())
				continue
			}
			if tick > scout.Start+s.MaxScoutTime {
				s.toReturn = append(s.toReturn, query.Entity())
				continue
			}
		}

		pos.X += dir.X * maxSpeed
		pos.Y += dir.Y * maxSpeed

		if pos.X < 0 || pos.X >= w {
			dir.X *= -1
			pos.X += dir.X * maxSpeed * 2
		}
		if pos.Y < 0 || pos.Y >= h {
			dir.Y *= -1
			pos.Y += dir.Y * maxSpeed * 2
		}
	}

	for _, e := range s.toForage {
		s.exchangeForage.Exchange(e)
		forage := s.forageMap.Get(e)
		forage.Start = tick
	}

	for _, e := range s.toReturn {
		s.exchangeReturn.Exchange(e)
	}

	s.toForage = s.toForage[:0]
	s.toReturn = s.toReturn[:0]
}

// Finalize the system
func (s *SysScouting) Finalize(world *ecs.World) {}
