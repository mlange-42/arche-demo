package evolution

import (
	"github.com/mlange-42/arche-model/resource"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// SysGrowGrassLinear is a system to grow the [Grass] resource.
type SysGrowGrassLinear struct {
	Interval int
	BaseRate float32

	grass generic.Resource[Grass]
	time  generic.Resource[resource.Tick]
}

// Initialize the system
func (s *SysGrowGrassLinear) Initialize(world *ecs.World) {
	s.grass = generic.NewResource[Grass](world)
	s.time = generic.NewResource[resource.Tick](world)
}

// Update the system
func (s *SysGrowGrassLinear) Update(world *ecs.World) {
	tick := s.time.Get().Tick
	if s.Interval > 0 && tick%int64(s.Interval) != 0 {
		return
	}

	grass := s.grass.Get()

	w, h := grass.Grass.Width(), grass.Grass.Height()
	f := s.BaseRate
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			v := grass.Grass.Get(x, y) + grass.Growth.Get(x, y)*f
			if v > 1 {
				v = 1
			}
			grass.Grass.Set(x, y, v)
		}
	}
}

// Finalize the system
func (s *SysGrowGrassLinear) Finalize(world *ecs.World) {}
