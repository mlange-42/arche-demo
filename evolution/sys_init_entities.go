package evolution

import (
	"math"
	"math/rand"

	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// SysInitEntities is a system to initialize entities.
type SysInitEntities struct {
	Count int
}

// Initialize the system
func (s *SysInitEntities) Initialize(world *ecs.World) {
	grassRes := generic.NewResource[Grass](world)
	grass := grassRes.Get()

	builder := generic.NewMap7[Position, Heading, Energy, Genotype, Phenotype, Color, Grazing](world)

	cs := float32(grass.Grass.Cellsize())
	w, h := float32(grass.Grass.Width())*cs, float32(grass.Grass.Height())*cs

	query := builder.NewBatchQ(s.Count)
	for query.Next() {
		pos, head, en, genes, pt, cols, _ := query.Get()

		pos.X = rand.Float32() * w
		pos.Y = rand.Float32() * h
		head.Angle = rand.Float32() * 2 * math.Pi

		genes.Randomize()
		pt.From(genes)

		cols.Color.R = uint8(50 + rand.Intn(200))
		cols.Color.G = uint8(50 + rand.Intn(200))
		cols.Color.B = uint8(50 + rand.Intn(200))
		cols.Color.A = 255

		en.Energy = 0.2 + rand.Float32()*0.8
	}
}

// Update the system
func (s *SysInitEntities) Update(world *ecs.World) {}

// Finalize the system
func (s *SysInitEntities) Finalize(world *ecs.World) {}
