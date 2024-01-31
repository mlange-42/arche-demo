package evolution

import (
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// SysDecisions is a system to perform grazer decisions.
type SysDecisions struct {
	grass          generic.Resource[Grass]
	grazerFilter   generic.Filter2[Position, Genes]
	searcherFilter generic.Filter2[Position, Genes]

	grazeExchange  generic.Exchange
	searchExchange generic.Exchange

	toGrass  []ecs.Entity
	toSearch []ecs.Entity
}

// Initialize the system
func (s *SysDecisions) Initialize(world *ecs.World) {
	s.grass = generic.NewResource[Grass](world)
	s.grazerFilter = *generic.NewFilter2[Position, Genes]().With(generic.T[Grazing]())
	s.searcherFilter = *generic.NewFilter2[Position, Genes]().With(generic.T[Searching]())

	s.grazeExchange = *generic.NewExchange(world).Adds(generic.T[Grazing]()).Removes(generic.T[Searching]())
	s.searchExchange = *generic.NewExchange(world).Adds(generic.T[Searching]()).Removes(generic.T[Grazing]())

	s.toGrass = make([]ecs.Entity, 0, 32)
	s.toSearch = make([]ecs.Entity, 0, 32)
}

// Update the system
func (s *SysDecisions) Update(world *ecs.World) {
	grass := &s.grass.Get().Grass

	queryG := s.grazerFilter.Query(world)
	for queryG.Next() {
		pos, genes := queryG.Get()
		cx, cy := grass.ToCell(float64(pos.X), float64(pos.Y))
		if grass.Get(cx, cy) < genes.MinGrass {
			s.toSearch = append(s.toSearch, queryG.Entity())
		}
	}

	queryS := s.searcherFilter.Query(world)
	for queryS.Next() {
		pos, genes := queryS.Get()
		cx, cy := grass.ToCell(float64(pos.X), float64(pos.Y))
		if grass.Get(cx, cy) > genes.MinGrass {
			s.toGrass = append(s.toGrass, queryS.Entity())
		}
	}

	for _, e := range s.toSearch {
		s.searchExchange.Exchange(e)
	}
	for _, e := range s.toGrass {
		s.grazeExchange.Exchange(e)
	}

	s.toGrass = s.toGrass[:0]
	s.toSearch = s.toSearch[:0]
}

// Finalize the system
func (s *SysDecisions) Finalize(world *ecs.World) {}
