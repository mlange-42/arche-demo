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
	MatingTrials  int
	MaxMatingDiff uint8
	CrossProb     float32
	AllowAsexual  bool

	filter    generic.Filter2[Energy, Color]
	parentMap generic.Map4[Position, Energy, Genes, Color]
	childMap  generic.Map6[Position, Heading, Energy, Genes, Color, Searching]
	mateMap   generic.Map2[Genes, Color]

	toReproduce []repEntry
}

// Initialize the system
func (s *SysReproduction) Initialize(world *ecs.World) {
	s.filter = *generic.NewFilter2[Energy, Color]()
	s.parentMap = generic.NewMap4[Position, Energy, Genes, Color](world)
	s.childMap = generic.NewMap6[Position, Heading, Energy, Genes, Color, Searching](world)
	s.mateMap = generic.NewMap2[Genes, Color](world)

	s.toReproduce = make([]repEntry, 0, 16)
}

// Update the system
func (s *SysReproduction) Update(world *ecs.World) {
	query := s.filter.Query(world)
	for query.Next() {
		en, col := query.Get()
		if en.Energy >= 1 {
			s.toReproduce = append(s.toReproduce, repEntry{Entity: query.Entity(), color: col.Color})
		}
	}

	for _, e := range s.toReproduce {
		pos, en, genes, col := s.parentMap.Get(e.Entity)

		enTotal := genes.Invest
		en.Energy -= enTotal
		enChild := enTotal / float32(genes.Offspring)

		mate, ok := s.findMate(&col.Color)
		if !(ok || s.AllowAsexual) {
			continue
		}

		query := s.childMap.NewBatchQ(int(genes.Offspring))
		for query.Next() {
			pos2, head2, en2, genes2, col2, _ := query.Get()
			head2.Angle = rand.Float32() * 2 * math.Pi
			en2.Energy = enChild
			*pos2 = *pos

			if ok {
				genesMate, colMate := s.mateMap.Get(mate)
				s.crossGenes(genes, genesMate, genes2)
				s.crossColor(col, colMate, col2)
			} else {
				*genes2 = *genes
				*col2 = *col
			}

			s.mutate(genes2, col2)
		}
	}

	s.toReproduce = s.toReproduce[:0]
}

// Finalize the system
func (s *SysReproduction) Finalize(world *ecs.World) {}

func (s *SysReproduction) crossGenes(g1, g2, result *Genes) {
	p := 1.0 - s.CrossProb
	if rand.Float32() < p {
		result.MaxAngle = g1.MaxAngle
	} else {
		result.MaxAngle = rand.Float32()*(g2.MaxAngle-g1.MaxAngle) + g1.MaxAngle
	}
	if rand.Float32() < p {
		result.MinGrass = g1.MinGrass
	} else {
		result.MinGrass = rand.Float32()*(g2.MinGrass-g1.MinGrass) + g1.MinGrass
	}
	if rand.Float32() < p {
		result.Invest = g1.Invest
	} else {
		result.Invest = rand.Float32()*(g2.Invest-g1.Invest) + g1.Invest
	}
	if rand.Float32() < p {
		result.Offspring = g1.Offspring
	} else {
		result.Offspring = common.RandBetweenUIn8(g1.Offspring, g2.Offspring)
	}
}

func (s *SysReproduction) crossColor(g1, g2, result *Color) {
	result.Color.R = common.RandBetweenUIn8(g1.Color.R, g2.Color.R)
	result.Color.G = common.RandBetweenUIn8(g1.Color.G, g2.Color.G)
	result.Color.B = common.RandBetweenUIn8(g1.Color.B, g2.Color.B)
}

func (s *SysReproduction) mutate(genes *Genes, color *Color) {
	genes.MaxAngle = common.Clamp32(common.RadToDeg*genes.MaxAngle+float32(rand.NormFloat64()*3), 10, 90) * common.DegToRad
	genes.MinGrass = common.Clamp32(genes.MinGrass+float32(rand.NormFloat64()*0.02), 0, 1)
	genes.Invest = common.Clamp32(genes.Invest+float32(rand.NormFloat64()*0.02), 0, 1)
	genes.Offspring = uint8(common.ClampInt(int(genes.Offspring)+rand.Intn(3)-1, 1, 10))

	color.Color.R = uint8(common.ClampInt(int(color.Color.R)+rand.Intn(5)-2, 50, 250))
	color.Color.G = uint8(common.ClampInt(int(color.Color.G)+rand.Intn(5)-2, 50, 250))
	color.Color.B = uint8(common.ClampInt(int(color.Color.B)+rand.Intn(5)-2, 50, 250))
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
