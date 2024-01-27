package bees

import (
	"math"
	"math/rand"

	"github.com/mlange-42/arche-demo/common"
	"github.com/mlange-42/arche-model/resource"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// SysScouting is a system that handles movement of scouting bees ([ActScout]).
//
// Switches activity to [ActReturn] after a certain time of scouting.
type SysScouting struct {
	MaxRotation  float64
	MaxScoutTime int64
	canvas       generic.Resource[common.Image]
	patches      generic.Resource[Patches]
	params       generic.Resource[Params]
	time         generic.Resource[resource.Tick]
	filter       generic.Filter4[Position, Direction, ActScout, Random256]

	exchangeForage generic.Exchange
	exchangeReturn generic.Exchange
	posMap         generic.Map1[Position]
	homeMap        generic.Map[HomeHive]
	forageMap      generic.Map1[ActForage]
	returnMap      generic.Map2[Position, ActReturn]

	toForage []ecs.Entity
	toReturn []ecs.Entity
}

// Initialize the system
func (s *SysScouting) Initialize(world *ecs.World) {
	s.canvas = generic.NewResource[common.Image](world)
	s.patches = generic.NewResource[Patches](world)
	s.params = generic.NewResource[Params](world)
	s.time = generic.NewResource[resource.Tick](world)

	s.filter = *generic.NewFilter4[Position, Direction, ActScout, Random256]()

	s.exchangeForage = *generic.NewExchange(world).
		Adds(generic.T[ActForage]()).
		Removes(generic.T[ActScout]())
	s.exchangeReturn = *generic.NewExchange(world).
		Adds(generic.T[ActReturn]()).
		Removes(generic.T[ActScout]())

	s.posMap = generic.NewMap1[Position](world)
	s.homeMap = generic.NewMap[HomeHive](world)
	s.forageMap = generic.NewMap1[ActForage](world)
	s.returnMap = generic.NewMap2[Position, ActReturn](world)

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
			dir.X, dir.Y = common.Rotate(dir.X, dir.Y, rand.Float64()*2*maxAng-maxAng)

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
		pos, ret := s.returnMap.Get(e)
		home := s.homeMap.GetRelation(e)
		hPos := s.posMap.Get(home)
		ret.Target = *hPos
		ret.Source = *pos
		ret.Load = 0
	}

	s.toForage = s.toForage[:0]
	s.toReturn = s.toReturn[:0]
}

// Finalize the system
func (s *SysScouting) Finalize(world *ecs.World) {}
