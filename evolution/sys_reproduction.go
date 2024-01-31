package evolution

import (
	"image/color"
	"math"
	"math/rand"

	"github.com/mlange-42/arche-demo/common"
	"github.com/mlange-42/arche-model/resource"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

type repEntry struct {
	Entity ecs.Entity
	color  color.RGBA
}

// SysReproduction is a system that handles reproduction of grazers.
type SysReproduction struct {
	MatingTrials        int
	MaxMatingDiff       uint8
	CrossProb           float32
	MutationProbability float32
	MutationMagnitude   float32
	AllowAsexual        bool
	HatchRadius         float32

	time      generic.Resource[resource.Tick]
	grass     generic.Resource[Grass]
	filter    generic.Filter2[Energy, Color]
	parentMap generic.Map5[Position, Energy, Genotype, Phenotype, Color]
	childMap  generic.Map8[Position, Age, Heading, Energy, Genotype, Phenotype, Color, Searching]
	mateMap   generic.Map1[Genotype]

	toReproduce []repEntry
}

// Initialize the system
func (s *SysReproduction) Initialize(world *ecs.World) {
	s.time = generic.NewResource[resource.Tick](world)
	s.grass = generic.NewResource[Grass](world)
	s.filter = *generic.NewFilter2[Energy, Color]()
	s.parentMap = generic.NewMap5[Position, Energy, Genotype, Phenotype, Color](world)
	s.childMap = generic.NewMap8[Position, Age, Heading, Energy, Genotype, Phenotype, Color, Searching](world)
	s.mateMap = generic.NewMap1[Genotype](world)

	s.toReproduce = make([]repEntry, 0, 16)
}

// Update the system
func (s *SysReproduction) Update(world *ecs.World) {
	tick := s.time.Get().Tick
	grass := s.grass.Get().Grass

	query := s.filter.Query(world)
	for query.Next() {
		en, col := query.Get()
		if en.Energy >= 1 {
			s.toReproduce = append(s.toReproduce, repEntry{Entity: query.Entity(), color: col.Color})
		}
	}

	for _, e := range s.toReproduce {
		pos, en, genes, pt, col := s.parentMap.Get(e.Entity)

		enTotal := pt.Invest * en.Energy
		en.Energy -= enTotal
		enChild := enTotal / float32(pt.Offspring)

		mate, ok := s.findMate(&col.Color)
		if !(ok || s.AllowAsexual) {
			continue
		}

		query := s.childMap.NewBatchQ(int(pt.Offspring))
		for query.Next() {
			pos2, age2, head2, en2, genes2, pt2, col2, _ := query.Get()
			head2.Angle = rand.Float32() * 2 * math.Pi
			en2.Energy = enChild
			age2.TickOfBirth = tick
			for {
				pos2.X = pos.X + float32(rand.NormFloat64())*s.HatchRadius
				pos2.Y = pos.Y + float32(rand.NormFloat64())*s.HatchRadius
				if grass.IsInBounds(float64(pos2.X), float64(pos2.Y)) {
					break
				}
			}
			if ok {
				genesMate := s.mateMap.Get(mate)
				s.cross(genes, genesMate, genes2)
			} else {
				*genes2 = *genes
			}

			s.mutate(genes2, col2)
			pt2.From(genes2)
			col2.From(genes2)
		}
	}

	s.toReproduce = s.toReproduce[:0]
}

// Finalize the system
func (s *SysReproduction) Finalize(world *ecs.World) {}

func (s *SysReproduction) cross(g1, g2, result *Genotype) {
	p := 1.0 - s.CrossProb
	result.Genes = g1.Genes
	for i := range result.Genes {
		if rand.Float32() < p {
			result.Genes[i] = rand.Float32()*(g2.Genes[i]-g1.Genes[i]) + g1.Genes[i]
		}
	}
}

func (s *SysReproduction) mutate(genes *Genotype, color *Color) {
	mag := s.MutationMagnitude
	for i := range genes.Genes {
		if rand.Float32() < s.MutationProbability {
			genes.Genes[i] = common.Clamp32(genes.Genes[i]+float32(rand.NormFloat64())*mag, 0, 1)
		}
	}
}

func (s *SysReproduction) findMate(col *color.RGBA) (ecs.Entity, bool) {
	ln := len(s.toReproduce)
	for i := 0; i < s.MatingTrials; i++ {
		entry := &s.toReproduce[rand.Intn(ln)]
		if s.canMate(col, &entry.color) {
			return entry.Entity, true
		}
	}

	return ecs.Entity{}, false
}

func (s *SysReproduction) canMate(a, b *color.RGBA) bool {
	return common.AbsInt(int(a.R)-int(b.R)) <= int(s.MaxMatingDiff) &&
		common.AbsInt(int(a.G)-int(b.G)) <= int(s.MaxMatingDiff) &&
		common.AbsInt(int(a.B)-int(b.B)) <= int(s.MaxMatingDiff)
}
