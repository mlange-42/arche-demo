package bees

import (
	"math"
	"math/rand"

	"github.com/mlange-42/arche-demo/common"
	"github.com/mlange-42/arche-model/resource"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// SysReturning is a system that handles movement of bees returning to their hive ([ActReturn]).
//
// Switches activity to [ActInHive] or [ActWaggleDance] on arrival.
// The probability of dancing depends on the resource load the bee brought back.
type SysReturning struct {
	MaxRotation         float64
	FleeDistance        float64
	MaxDanceProbability float64

	params         generic.Resource[Params]
	time           generic.Resource[resource.Tick]
	mouse          generic.Resource[common.Mouse]
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
	s.mouse = generic.NewResource[common.Mouse](world)
	s.filter = *generic.NewFilter4[Position, Direction, ActReturn, Random256]()

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
	maxAng := (s.MaxRotation * math.Pi / 180.0) / 2

	mouse := s.mouse.Get()
	mouseInside := mouse.IsInside

	fleeDistSq := s.FleeDistance * s.FleeDistance

	query := s.filter.Query(world)
	for query.Next() {
		pos, dir, ret, r256 := query.Get()

		speed := 1.0
		if tick%4 == int64(r256.Value)%4 {
			var dx, dy float64
			if mouseInside && common.DistanceSq(pos.X, pos.Y, float64(mouse.X), float64(mouse.Y)) < fleeDistSq {
				dx = pos.X - float64(mouse.X)
				dy = pos.Y - float64(mouse.Y)
				speed = 1.5
			} else {
				dx = ret.Target.X - pos.X
				dy = ret.Target.Y - pos.Y
			}

			dir.X, dir.Y, _ = common.Norm(dx, dy)
			dir.X, dir.Y = common.Rotate(dir.X, dir.Y, rand.Float64()*2*maxAng-maxAng)

			if dx*dx+dy*dy < 4 {
				s.toArrive = append(s.toArrive, query.Entity())
				continue
			}
		}

		pos.X += dir.X * speed * maxSpeed
		pos.Y += dir.Y * speed * maxSpeed
	}

	for _, e := range s.toArrive {
		ret := s.returnMap.Get(e)
		if ret.Load <= 0 || rand.Float64() >= ret.Load*s.MaxDanceProbability {
			s.exchangeArrive.Exchange(e)
			continue
		}
		trg := ret.Source
		dist := common.Distance(trg.X, trg.Y, ret.Target.X, ret.Target.Y)
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
