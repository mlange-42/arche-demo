package bees

import (
	"github.com/mlange-42/arche-model/resource"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// SysForaging is a system that handles resource extraction of foraging bees ([ActForage]) from patches.
//
// Switches activity to [ActReturn] after a certain foraging time.
type SysForaging struct {
	MaxForagingTime int64
	MaxCollect      float64
	patches         generic.Resource[Patches]
	time            generic.Resource[resource.Tick]
	filter          generic.Filter2[Position, ActForage]

	exchangeReturn generic.Exchange

	posMap    generic.Map1[Position]
	homeMap   generic.Map[HomeHive]
	returnMap generic.Map2[Position, ActReturn]
	patchMap  generic.Map[FlowerPatch]
	forageMap generic.Map1[ActForage]

	toReturn []ecs.Entity
}

// Initialize the system
func (s *SysForaging) Initialize(world *ecs.World) {
	s.patches = generic.NewResource[Patches](world)
	s.time = generic.NewResource[resource.Tick](world)

	s.filter = *generic.NewFilter2[Position, ActForage]()

	s.exchangeReturn = *generic.NewExchange(world).
		Adds(generic.T[ActReturn]()).
		Removes(generic.T[ActForage]())

	s.posMap = generic.NewMap1[Position](world)
	s.homeMap = generic.NewMap[HomeHive](world)
	s.patchMap = generic.NewMap[FlowerPatch](world)
	s.returnMap = generic.NewMap2[Position, ActReturn](world)
	s.forageMap = generic.NewMap1[ActForage](world)

	s.toReturn = make([]ecs.Entity, 0, 64)
}

// Update the system
func (s *SysForaging) Update(world *ecs.World) {
	patches := s.patches.Get()
	tick := s.time.Get().Tick

	query := s.filter.Query(world)
	for query.Next() {
		pos, forage := query.Get()

		if forage.Load <= 0 {
			x, y := patches.ToCell(pos.X, pos.Y)
			patchEntity := patches.Patches[x][y]
			if patchEntity.IsZero() {
				s.toReturn = append(s.toReturn, query.Entity())
				continue
			}
			patch := s.patchMap.Get(patchEntity)
			if patch.Resources <= 0 {
				s.toReturn = append(s.toReturn, query.Entity())
				continue
			}
			forage.Load = patch.Resources
			if patch.Resources > s.MaxCollect {
				patch.Resources -= s.MaxCollect
			} else {
				patch.Resources = 0
			}
		}

		if tick > forage.Start+s.MaxForagingTime {
			s.toReturn = append(s.toReturn, query.Entity())
			continue
		}
	}

	for _, e := range s.toReturn {
		load := s.forageMap.Get(e).Load

		s.exchangeReturn.Exchange(e)
		pos, ret := s.returnMap.Get(e)
		home := s.homeMap.GetRelation(e)
		hPos := s.posMap.Get(home)

		ret.Target = *hPos

		sx, sy := patches.ToCellCenter(pos.X, pos.Y)
		ret.Source = Position{X: sx, Y: sy}
		ret.Load = load
	}

	s.toReturn = s.toReturn[:0]
}

// Finalize the system
func (s *SysForaging) Finalize(world *ecs.World) {}
