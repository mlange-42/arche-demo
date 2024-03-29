package bees

import (
	"math/rand"

	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// SysManagePatches is a system that creates a number of randomly placed flower patches.
// Further, it removes patches that are depleted, and creates new patches if the
// number of patches falls below Count.
type SysManagePatches struct {
	// Target number of patches.
	Count int

	patchFilter  generic.Filter1[FlowerPatch]
	patchBuilder generic.Map1[FlowerPatch]
	patchesRes   generic.Resource[Patches]
	patchMap     generic.Map[FlowerPatch]

	toRemove []ecs.Entity
}

// Initialize the system
func (s *SysManagePatches) Initialize(world *ecs.World) {
	s.patchFilter = *generic.NewFilter1[FlowerPatch]()
	s.patchBuilder = generic.NewMap1[FlowerPatch](world)
	s.patchesRes = generic.NewResource[Patches](world)
	s.patchMap = generic.NewMap[FlowerPatch](world)
	s.toRemove = make([]ecs.Entity, 0, 16)

	s.createRandomPatches(s.Count)
}

// Update the system
func (s *SysManagePatches) Update(world *ecs.World) {
	patches := s.patchesRes.Get()

	query := s.patchFilter.Query(world)
	count := 0

	for query.Next() {
		patch := query.Get()

		if patch.Resources <= 0 {
			s.toRemove = append(s.toRemove, query.Entity())
		} else {
			count++
		}
	}

	for _, e := range s.toRemove {
		patch := s.patchMap.Get(e)
		patches.Patches[patch.X][patch.Y] = ecs.Entity{}
		world.RemoveEntity(e)
	}

	if count < s.Count {
		s.createRandomPatches(s.Count - count)
	}

	s.toRemove = s.toRemove[:0]
}

// Finalize the system
func (s *SysManagePatches) Finalize(world *ecs.World) {}

func (s *SysManagePatches) createRandomPatches(count int) {
	patches := s.patchesRes.Get()
	w := patches.Cols
	h := patches.Rows

	query := s.patchBuilder.NewBatchQ(count)
	for query.Next() {
		e := query.Entity()
		patch := query.Get()

		for {
			x := rand.Intn(w)
			y := rand.Intn(h)
			if patches.Patches[x][y].IsZero() {
				patches.Patches[x][y] = e
				patch.X = x
				patch.Y = y
				patch.Resources = 1.0
				break
			}
		}
	}
}
