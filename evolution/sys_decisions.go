package evolution

import (
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// SysDecisions is a system to perform grazer decisions.
type SysDecisions struct {
	grass  generic.Resource[Grass]
	filter generic.Filter3[Position, Phenotype, Activity]
}

// Initialize the system
func (s *SysDecisions) Initialize(world *ecs.World) {
	s.grass = generic.NewResource[Grass](world)
	s.filter = *generic.NewFilter3[Position, Phenotype, Activity]()
}

// Update the system
func (s *SysDecisions) Update(world *ecs.World) {
	grass := &s.grass.Get().Grass

	query := s.filter.Query(world)
	for query.Next() {
		pos, pt, act := query.Get()
		cx, cy := grass.ToCell(float64(pos.X), float64(pos.Y))
		if act.IsGrazing {
			if grass.Get(cx, cy) < pt.MinGrass {
				act.IsGrazing = false
			}
		} else {
			if grass.Get(cx, cy) > pt.MinGrass {
				act.IsGrazing = true
			}
		}
	}
}

// Finalize the system
func (s *SysDecisions) Finalize(world *ecs.World) {}
