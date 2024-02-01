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
	Pos    Position
	Color  color.RGBA
}

// SysReproduction is a system that handles reproduction of grazers.
type SysReproduction struct {
	MatingTrials        int
	MaxMatingDist       int
	MaxMatingDiff       uint8
	CrossProb           float32
	MutationProbability float32
	MutationMagnitude   float32
	AllowAsexual        bool
	HatchRadius         float32

	maxMatingDistSq int

	time      generic.Resource[resource.Tick]
	grass     generic.Resource[Grass]
	filter    generic.Filter3[Position, Energy, Color]
	parentMap generic.Map5[Position, Energy, Genotype, Phenotype, Color]
	childMap  generic.Map8[Position, Age, Heading, Energy, Genotype, Phenotype, Color, Activity]
	mateMap   generic.Map1[Genotype]

	toReproduce []repEntry
}

// Initialize the system
func (s *SysReproduction) Initialize(world *ecs.World) {
	s.time = generic.NewResource[resource.Tick](world)
	s.grass = generic.NewResource[Grass](world)
	s.filter = *generic.NewFilter3[Position, Energy, Color]()
	s.parentMap = generic.NewMap5[Position, Energy, Genotype, Phenotype, Color](world)
	s.childMap = generic.NewMap8[Position, Age, Heading, Energy, Genotype, Phenotype, Color, Activity](world)
	s.mateMap = generic.NewMap1[Genotype](world)

	s.maxMatingDistSq = s.MaxMatingDist * s.MaxMatingDist

	s.toReproduce = make([]repEntry, 0, 16)
}

// Update the system
func (s *SysReproduction) Update(world *ecs.World) {
	tick := s.time.Get().Tick
	grass := s.grass.Get().Grass

	query := s.filter.Query(world)
	for query.Next() {
		pos, en, col := query.Get()
		if en.Energy >= 1 {
			s.toReproduce = append(s.toReproduce, repEntry{Entity: query.Entity(), Pos: *pos, Color: col.Color})
		}
	}

	for _, e := range s.toReproduce {
		pos, en, genes, pt, col := s.parentMap.Get(e.Entity)

		mate, ok := s.findMate(e.Entity, pos, &col.Color)
		var genesMate *Genotype
		if !(ok || s.AllowAsexual) {
			continue
		}
		if ok {
			genesMate = s.mateMap.Get(mate)
		}

		enTotal := pt.Invest * en.Energy
		en.Energy -= enTotal
		enChild := enTotal / float32(pt.Offspring)

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

func (s *SysReproduction) findMate(e ecs.Entity, pos *Position, col *color.RGBA) (ecs.Entity, bool) {
	ln := len(s.toReproduce)
	for i := 0; i < s.MatingTrials; i++ {
		entry := &s.toReproduce[rand.Intn(ln)]
		if e != entry.Entity && s.canMate(pos, col, entry) {
			return entry.Entity, true
		}
	}

	return ecs.Entity{}, false
}

func (s *SysReproduction) canMate(pos *Position, col *color.RGBA, other *repEntry) bool {
	md := int(s.MaxMatingDiff)
	if common.AbsInt(int(col.R)-int(other.Color.R)) > md ||
		common.AbsInt(int(col.G)-int(other.Color.G)) > md ||
		common.AbsInt(int(col.B)-int(other.Color.B)) > md {
		return false
	}
	if s.MaxMatingDist > 0 {
		dx := pos.X - other.Pos.X
		dy := pos.Y - other.Pos.Y
		if dx*dx+dy*dy > float32(s.maxMatingDistSq) {
			return false
		}
	}
	return true
}
