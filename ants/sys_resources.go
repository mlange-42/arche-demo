package ants

import (
	"math/rand"

	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// SysResources is a system that initializes and manages collectible resources.
type SysResources struct {
	Count int

	patches         generic.Resource[Patches]
	resourcesFilter generic.Filter1[NodeResource]
	resourceMap     generic.Map[NodeResource]
	resourceAdder   generic.Map1[NodeResource]

	exchangeRemove generic.Exchange

	toRemove []ecs.Entity
}

// Initialize the system
func (s *SysResources) Initialize(world *ecs.World) {
	s.patches = generic.NewResource[Patches](world)
	s.resourcesFilter = *generic.NewFilter1[NodeResource]()
	s.resourceMap = generic.NewMap[NodeResource](world)
	s.resourceAdder = generic.NewMap1[NodeResource](world)

	s.exchangeRemove = *generic.NewExchange(world).Removes(generic.T[NodeResource]())
	s.toRemove = make([]ecs.Entity, 0, 8)

	s.createRandomResources(s.Count)
}

// Update the system
func (s *SysResources) Update(world *ecs.World) {
	query := s.resourcesFilter.Query(world)
	count := 0
	for query.Next() {
		res := query.Get()

		if res.Resource <= 0 {
			s.toRemove = append(s.toRemove, query.Entity())
		} else {
			count++
		}
	}

	for _, e := range s.toRemove {
		s.exchangeRemove.Exchange(e)
	}

	if count < s.Count {
		s.createRandomResources(s.Count - count)
	}

	s.toRemove = s.toRemove[:0]
}

// Finalize the system
func (s *SysResources) Finalize(world *ecs.World) {}

func (s *SysResources) createRandomResources(count int) {
	patches := s.patches.Get()

	w := patches.Cols
	h := patches.Rows

	for i := 0; i < count; i++ {
		for true {
			x := rand.Intn(w)
			y := rand.Intn(h)
			node := patches.Get(x, y)
			if !s.resourceMap.Has(node) {
				s.resourceAdder.Assign(node, &NodeResource{Resource: 1.0})
				break
			}
		}
	}
}
