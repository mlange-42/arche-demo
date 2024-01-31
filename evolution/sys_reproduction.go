package evolution

import (
	"math"
	"math/rand"

	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// SysReproduction is a system that handles reproduction of grazers.
type SysReproduction struct {
	filter    generic.Filter1[Energy]
	parentMap generic.Map4[Position, Energy, Genes, Color]
	childMap  generic.Map6[Position, Heading, Energy, Genes, Color, Searching]

	toReproduce []ecs.Entity
}

// Initialize the system
func (s *SysReproduction) Initialize(world *ecs.World) {
	s.filter = *generic.NewFilter1[Energy]()
	s.parentMap = generic.NewMap4[Position, Energy, Genes, Color](world)
	s.childMap = generic.NewMap6[Position, Heading, Energy, Genes, Color, Searching](world)

	s.toReproduce = make([]ecs.Entity, 0, 16)
}

// Update the system
func (s *SysReproduction) Update(world *ecs.World) {
	query := s.filter.Query(world)
	for query.Next() {
		en := query.Get()
		if en.Energy >= 1 {
			s.toReproduce = append(s.toReproduce, query.Entity())
		}
	}

	for _, e := range s.toReproduce {
		pos, en, genes, col := s.parentMap.Get(e)
		enTotal := genes.Invest
		en.Energy -= enTotal
		enChild := enTotal / float32(genes.Offspring)

		query := s.childMap.NewBatchQ(int(genes.Offspring))
		for query.Next() {
			pos2, head2, en2, genes2, col2, _ := query.Get()

			*pos2 = *pos
			head2.Angle = rand.Float32() * 2 * math.Pi

			*genes2 = *genes
			*col2 = *col

			en2.Energy = enChild
		}
	}

	s.toReproduce = s.toReproduce[:0]
}

// Finalize the system
func (s *SysReproduction) Finalize(world *ecs.World) {}
