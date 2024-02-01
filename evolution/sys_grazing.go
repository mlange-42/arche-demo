package evolution

import (
	"github.com/mlange-42/arche-demo/common"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// SysGrazing is a system that handle grazing of entities on [Grass].
type SysGrazing struct {
	MaxUptake    float32
	UptakeFactor float32

	grass        generic.Resource[Grass]
	grazerFilter generic.Filter2[Position, Activity]
	energyFilter generic.Filter3[Position, Activity, Energy]

	grazers common.Grid[int16]
}

// Initialize the system
func (s *SysGrazing) Initialize(world *ecs.World) {
	s.grass = generic.NewResource[Grass](world)
	s.grazerFilter = *generic.NewFilter2[Position, Activity]()
	s.energyFilter = *generic.NewFilter3[Position, Activity, Energy]()

	grass := &s.grass.Get().Grass
	s.grazers = common.NewGrid[int16](grass.Width(), grass.Height(), grass.Cellsize())
}

// Update the system
func (s *SysGrazing) Update(world *ecs.World) {
	grass := &s.grass.Get().Grass

	s.grazers.Fill(0)

	query := s.grazerFilter.Query(world)
	for query.Next() {
		pos, act := query.Get()
		if !act.IsGrazing {
			continue
		}
		cx, cy := s.grazers.ToCell(float64(pos.X), float64(pos.Y))
		*s.grazers.GetPointer(cx, cy)++
	}

	queryEn := s.energyFilter.Query(world)
	for queryEn.Next() {
		pos, act, en := queryEn.Get()
		if !act.IsGrazing {
			continue
		}
		cx, cy := s.grazers.ToCell(float64(pos.X), float64(pos.Y))
		available := grass.Get(cx, cy) / float32(s.grazers.Get(cx, cy))
		uptake := grass.Get(cx, cy) * s.MaxUptake
		if uptake > available {
			uptake = available
		}
		en.Energy += uptake * s.UptakeFactor
		if en.Energy > 1 {
			en.Energy = 1
		}
		*grass.GetPointer(cx, cy) -= uptake
	}
}

// Finalize the system
func (s *SysGrazing) Finalize(world *ecs.World) {}
