package evolution

import (
	"image/color"
	"math"
	"math/rand"

	"github.com/mlange-42/arche-demo/common"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

type repEntry struct {
	Entity ecs.Entity
	color  color.RGBA
}

// SysReproduction is a system that handles reproduction of grazers.
type SysReproduction struct {
	MatingTrials           int
	MaxMatingDiff          uint8
	CrossProb              float32
	MutationMagnitude      float32
	ColorMutationMagnitude uint8
	AllowAsexual           bool
	HatchRadius            float32

	grass     generic.Resource[Grass]
	filter    generic.Filter2[Energy, Color]
	parentMap generic.Map5[Position, Energy, Genotype, Phenotype, Color]
	childMap  generic.Map7[Position, Heading, Energy, Genotype, Phenotype, Color, Searching]
	mateMap   generic.Map2[Genotype, Color]

	toReproduce []repEntry
}

// Initialize the system
func (s *SysReproduction) Initialize(world *ecs.World) {
	s.grass = generic.NewResource[Grass](world)
	s.filter = *generic.NewFilter2[Energy, Color]()
	s.parentMap = generic.NewMap5[Position, Energy, Genotype, Phenotype, Color](world)
	s.childMap = generic.NewMap7[Position, Heading, Energy, Genotype, Phenotype, Color, Searching](world)
	s.mateMap = generic.NewMap2[Genotype, Color](world)

	s.toReproduce = make([]repEntry, 0, 16)
}

// Update the system
func (s *SysReproduction) Update(world *ecs.World) {
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

		enTotal := pt.Invest
		en.Energy -= enTotal
		enChild := enTotal / float32(pt.Offspring)

		mate, ok := s.findMate(&col.Color)
		if !(ok || s.AllowAsexual) {
			continue
		}

		query := s.childMap.NewBatchQ(int(pt.Offspring))
		for query.Next() {
			pos2, head2, en2, genes2, pt2, col2, _ := query.Get()
			head2.Angle = rand.Float32() * 2 * math.Pi
			en2.Energy = enChild
			for {
				pos2.X = pos.X + float32(rand.NormFloat64())*s.HatchRadius
				pos2.Y = pos.Y + float32(rand.NormFloat64())*s.HatchRadius
				if grass.IsInBounds(float64(pos2.X), float64(pos2.Y)) {
					break
				}
			}
			if ok {
				genesMate, colMate := s.mateMap.Get(mate)
				s.crossGenes(genes, genesMate, genes2)
				s.crossColor(col, colMate, col2)
			} else {
				*genes2 = *genes
				*col2 = *col
			}

			s.mutate(genes2, col2)
			pt2.From(genes2)
		}
	}

	s.toReproduce = s.toReproduce[:0]
}

// Finalize the system
func (s *SysReproduction) Finalize(world *ecs.World) {}

func (s *SysReproduction) crossGenes(g1, g2, result *Genotype) {
	p := 1.0 - s.CrossProb
	result.Genes = g1.Genes
	for i := range result.Genes {
		if rand.Float32() < p {
			result.Genes[i] = rand.Float32()*(g2.Genes[i]-g1.Genes[i]) + g1.Genes[i]
		}
	}
}

func (s *SysReproduction) crossColor(g1, g2, result *Color) {
	p := 1.0 - s.CrossProb
	result.Color = g1.Color
	if rand.Float32() < p {
		result.Color.R = common.RandBetweenUIn8(g1.Color.R, g2.Color.R)
	}
	if rand.Float32() < p {
		result.Color.G = common.RandBetweenUIn8(g1.Color.G, g2.Color.G)
	}
	if rand.Float32() < p {
		result.Color.B = common.RandBetweenUIn8(g1.Color.B, g2.Color.B)
	}
}

func (s *SysReproduction) mutate(genes *Genotype, color *Color) {
	mag := s.MutationMagnitude
	for i := range genes.Genes {
		genes.Genes[i] = common.Clamp32(genes.Genes[i]+float32(rand.NormFloat64())*mag, 0, 1)
	}

	cmHalf := int(s.ColorMutationMagnitude)
	cm := cmHalf*2 + 1
	color.Color.R = uint8(common.ClampInt(int(color.Color.R)+rand.Intn(cm)-cmHalf, 50, 250))
	color.Color.G = uint8(common.ClampInt(int(color.Color.G)+rand.Intn(cm)-cmHalf, 50, 250))
	color.Color.B = uint8(common.ClampInt(int(color.Color.B)+rand.Intn(cm)-cmHalf, 50, 250))
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
