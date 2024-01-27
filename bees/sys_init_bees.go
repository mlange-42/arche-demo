package bees

import (
	"math/rand"

	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// InitBees is a system that creates a number of bee entities per hive.
type InitBees struct {
	// Target number of bees per hive.
	CountPerHive int
}

// Initialize the system
func (s *InitBees) Initialize(world *ecs.World) {
	builder := generic.NewMap5[Position, Direction, HomeHive, ActInHive, Random256](world, generic.T[HomeHive]())
	posMap := generic.NewMap1[Position](world)

	hives := make([]ecs.Entity, 0, 100)
	hiveFilter := generic.NewFilter0().With(generic.T[Hive]())
	query := hiveFilter.Query(world)
	for query.Next() {
		hives = append(hives, query.Entity())
	}

	for _, e := range hives {
		hivePos := posMap.Get(e)
		query := builder.NewBatchQ(s.CountPerHive, e)
		for query.Next() {
			pos, dir, _, _, r256 := query.Get()
			pos.X = hivePos.X
			pos.Y = hivePos.Y

			dir.X, dir.Y, _ = norm(rand.NormFloat64(), rand.NormFloat64())

			r256.Value = uint8(rand.Int31n(256))
		}
	}
}

// Update the system
func (s *InitBees) Update(world *ecs.World) {}

// Finalize the system
func (s *InitBees) Finalize(world *ecs.World) {}
