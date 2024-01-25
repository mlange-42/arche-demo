package main

import (
	"math/rand"

	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// InitPatches system
type InitPatches struct {
	Count int
}

// Initialize the system
func (s *InitPatches) Initialize(world *ecs.World) {
	patchesRes := generic.NewResource[Patches](world)
	patches := patchesRes.Get()
	w := patches.Cols
	h := patches.Rows

	builder := generic.NewMap1[FlowerPatch](world)

	query := builder.NewBatchQ(s.Count)
	for query.Next() {
		e := query.Entity()
		patch := query.Get()

		for true {
			x := rand.Intn(w)
			y := rand.Intn(h)
			if patches.Patches[x][y].IsZero() {
				patches.Patches[x][y] = e
				patch.X = x
				patch.Y = y
				break
			}
		}
	}
}

// Update the system
func (s *InitPatches) Update(world *ecs.World) {}

// Finalize the system
func (s *InitPatches) Finalize(world *ecs.World) {}
