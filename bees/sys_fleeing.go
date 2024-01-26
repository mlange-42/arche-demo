package main

import (
	"github.com/mlange-42/arche-demo/common"
	"github.com/mlange-42/arche-model/resource"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// SysFleeing system
type SysFleeing struct {
	FleeDistance float64
	time         generic.Resource[resource.Tick]
	mouse        generic.Resource[MouseListener]

	returnMap generic.Map1[ActReturn]
	posMap    generic.Map1[Position]
	homeMap   generic.Map[HomeHive]

	forageFilter   generic.Filter2[Position, Random256]
	forageExchange generic.Exchange

	scoutFilter   generic.Filter2[Position, Random256]
	scoutExchange generic.Exchange

	followFilter   generic.Filter2[Position, Random256]
	followExchange generic.Exchange

	toFlee []ecs.Entity
}

// Initialize the system
func (s *SysFleeing) Initialize(world *ecs.World) {
	s.time = generic.NewResource[resource.Tick](world)
	s.mouse = generic.NewResource[MouseListener](world)

	s.returnMap = generic.NewMap1[ActReturn](world)
	s.posMap = generic.NewMap1[Position](world)
	s.homeMap = generic.NewMap[HomeHive](world)

	s.forageFilter = *generic.NewFilter2[Position, Random256]().With(generic.T[ActForage]())
	s.forageExchange = *generic.NewExchange(world).Adds(generic.T[ActReturn]()).Removes(generic.T[ActForage]())

	s.scoutFilter = *generic.NewFilter2[Position, Random256]().With(generic.T[ActScout]())
	s.scoutExchange = *generic.NewExchange(world).Adds(generic.T[ActReturn]()).Removes(generic.T[ActScout]())

	s.followFilter = *generic.NewFilter2[Position, Random256]().With(generic.T[ActFollow]())
	s.followExchange = *generic.NewExchange(world).Adds(generic.T[ActReturn]()).Removes(generic.T[ActFollow]())

	s.toFlee = make([]ecs.Entity, 0, 64)
}

// Update the system
func (s *SysFleeing) Update(world *ecs.World) {
	tick := s.time.Get().Tick

	listener := s.mouse.Get()
	mouse := listener.Mouse
	mouseInside := listener.MouseInside
	if !mouseInside {
		return
	}

	s.fleeForage(world, mouse, tick)
	s.fleeScout(world, mouse, tick)
	s.fleeFollow(world, mouse, tick)
}

func (s *SysFleeing) fleeForage(world *ecs.World, mouse common.MousePointer, tick int64) {
	modTick := tick % 8
	fleeDistSq := s.FleeDistance * s.FleeDistance

	query := s.forageFilter.Query(world)
	for query.Next() {
		pos, r256 := query.Get()
		if modTick == int64(r256.Value)%8 && distanceSq(pos.X, pos.Y, mouse.X, mouse.Y) < fleeDistSq {
			s.toFlee = append(s.toFlee, query.Entity())
		}
	}
	s.exchange(&s.forageExchange, s.toFlee)
	s.toFlee = s.toFlee[:0]
}

func (s *SysFleeing) fleeScout(world *ecs.World, mouse common.MousePointer, tick int64) {
	modTick := tick % 8
	fleeDistSq := s.FleeDistance * s.FleeDistance

	query := s.scoutFilter.Query(world)
	for query.Next() {
		pos, r256 := query.Get()
		if modTick == int64(r256.Value)%8 && distanceSq(pos.X, pos.Y, mouse.X, mouse.Y) < fleeDistSq {
			s.toFlee = append(s.toFlee, query.Entity())
		}
	}
	s.exchange(&s.scoutExchange, s.toFlee)
	s.toFlee = s.toFlee[:0]
}

func (s *SysFleeing) fleeFollow(world *ecs.World, mouse common.MousePointer, tick int64) {
	modTick := tick % 8
	fleeDistSq := s.FleeDistance * s.FleeDistance

	query := s.followFilter.Query(world)
	for query.Next() {
		pos, r256 := query.Get()
		if modTick == int64(r256.Value)%8 && distanceSq(pos.X, pos.Y, mouse.X, mouse.Y) < fleeDistSq {
			s.toFlee = append(s.toFlee, query.Entity())
		}
	}
	s.exchange(&s.followExchange, s.toFlee)
	s.toFlee = s.toFlee[:0]
}

func (s *SysFleeing) exchange(ex *generic.Exchange, entities []ecs.Entity) {
	for _, e := range entities {
		ex.Exchange(e)
		ret := s.returnMap.Get(e)
		home := s.homeMap.GetRelation(e)
		hPos := s.posMap.Get(home)
		ret.Target = *hPos
	}
}

// Finalize the system
func (s *SysFleeing) Finalize(world *ecs.World) {}
