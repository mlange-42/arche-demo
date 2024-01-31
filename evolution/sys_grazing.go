package evolution

import (
	"math"

	"github.com/mlange-42/arche-demo/common"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// SysGrazing is a system that handle grazing of entities on [Grass].
type SysGrazing struct {
	MaxUptake float32

	grass        generic.Resource[Grass]
	grazerFilter generic.Filter1[Position]
	energyFilter generic.Filter2[Position, Energy]

	grazers common.Grid[int16]
}

// Initialize the system
func (s *SysGrazing) Initialize(world *ecs.World) {
	s.grass = generic.NewResource[Grass](world)
	s.grazerFilter = *generic.NewFilter1[Position]().With(generic.T[Grazing]())
	s.energyFilter = *generic.NewFilter2[Position, Energy]().With(generic.T[Grazing]())

	grass := &s.grass.Get().Grass
	s.grazers = common.NewGrid[int16](grass.Width(), grass.Height(), grass.Cellsize())
}

// Update the system
func (s *SysGrazing) Update(world *ecs.World) {
	grass := &s.grass.Get().Grass

	s.grazers.Fill(0)

	query := s.grazerFilter.Query(world)
	for query.Next() {
		pos := query.Get()
		cx, cy := s.grazers.ToCell(float64(pos.X), float64(pos.Y))
		*s.grazers.GetPointer(cx, cy)++
	}

	queryEn := s.energyFilter.Query(world)
	for queryEn.Next() {
		pos, en := queryEn.Get()
		cx, cy := s.grazers.ToCell(float64(pos.X), float64(pos.Y))
		uptake := float32(math.Pow(float64(grass.Get(cx, cy)), 2) / float64(s.grazers.Get(cx, cy)))
		if uptake > s.MaxUptake {
			uptake = s.MaxUptake
		}
		en.Energy += uptake
		*grass.GetPointer(cx, cy) -= uptake
	}
}

// Finalize the system
func (s *SysGrazing) Finalize(world *ecs.World) {}
