package evolution

import (
	"github.com/mlange-42/arche-model/resource"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// SysGrowGrassLogistic is a system to grow the [Grass] resource.
type SysGrowGrassLogistic struct {
	Interval int
	BaseRate float32

	grass generic.Resource[Grass]
	time  generic.Resource[resource.Tick]
}

// Initialize the system
func (s *SysGrowGrassLogistic) Initialize(world *ecs.World) {
	s.grass = generic.NewResource[Grass](world)
	s.time = generic.NewResource[resource.Tick](world)
}

// Update the system
func (s *SysGrowGrassLogistic) Update(world *ecs.World) {
	tick := s.time.Get().Tick
	if s.Interval > 0 && tick%int64(s.Interval) != 0 {
		return
	}

	grass := s.grass.Get()

	w, h := grass.Grass.Width(), grass.Grass.Height()
	f := s.BaseRate
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			gr := grass.Grass.Get(x, y)
			r := grass.Growth.Get(x, y) * f
			v := gr + r*gr*(1-gr)
			grass.Grass.Set(x, y, v)
		}
	}
}

// Finalize the system
func (s *SysGrowGrassLogistic) Finalize(world *ecs.World) {}
