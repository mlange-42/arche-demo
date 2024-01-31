package evolution

import (
	"math"
	"math/rand"

	"github.com/mlange-42/arche-model/resource"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// SysInitEntities is a system to initialize entities.
type SysInitEntities struct {
	InitialCount    int
	ReleaseInterval int
	ReleaseCount    int
	RandomGenes     bool

	width  float32
	height float32

	time    generic.Resource[resource.Tick]
	builder generic.Map7[Position, Heading, Energy, Genotype, Phenotype, Color, Grazing]
}

// Initialize the system
func (s *SysInitEntities) Initialize(world *ecs.World) {
	s.time = generic.NewResource[resource.Tick](world)
	s.builder = generic.NewMap7[Position, Heading, Energy, Genotype, Phenotype, Color, Grazing](world)

	grassRes := generic.NewResource[Grass](world)
	grass := grassRes.Get()

	cs := float32(grass.Grass.Cellsize())
	s.width, s.height = float32(grass.Grass.Width())*cs, float32(grass.Grass.Height())*cs

	s.createEntities(world, s.InitialCount)
}

// Update the system
func (s *SysInitEntities) Update(world *ecs.World) {
	tick := s.time.Get().Tick

	if s.ReleaseCount <= 0 || (s.ReleaseInterval > 0 && tick%int64(s.ReleaseInterval) != 0) {
		return
	}
	s.createEntities(world, s.ReleaseCount)
}

// Finalize the system
func (s *SysInitEntities) Finalize(world *ecs.World) {}

func (s *SysInitEntities) createEntities(world *ecs.World, count int) {
	query := s.builder.NewBatchQ(s.InitialCount)
	for query.Next() {
		pos, head, en, genes, pt, cols, _ := query.Get()

		pos.X = rand.Float32() * s.width
		pos.Y = rand.Float32() * s.height
		head.Angle = rand.Float32() * 2 * math.Pi

		if s.RandomGenes {
			genes.Randomize()
			cols.Randomize()
		} else {
			genes.Defaults()
			cols.Defaults()
		}
		pt.From(genes)

		en.Energy = 0.2 + rand.Float32()*0.8
	}
}
