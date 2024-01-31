package evolution

import (
	"math/rand"

	"github.com/mlange-42/arche-demo/common"
	"github.com/mlange-42/arche-model/resource"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// SysDisturbance is a system that applies disturbances.
type SysDisturbance struct {
	Interval    int
	Count       int
	MinRadius   int
	MaxRadius   int
	TargetValue float32

	grass generic.Resource[Grass]
	time  generic.Resource[resource.Tick]
}

// Initialize the system
func (s *SysDisturbance) Initialize(world *ecs.World) {
	s.grass = generic.NewResource[Grass](world)
	s.time = generic.NewResource[resource.Tick](world)
}

// Update the system
func (s *SysDisturbance) Update(world *ecs.World) {
	tick := s.time.Get().Tick
	if s.Interval > 0 && tick%int64(s.Interval) != 0 {
		return
	}

	grass := &s.grass.Get().Grass
	for i := 0; i < s.Count; i++ {
		x, y := rand.Intn(grass.Width()), rand.Intn(grass.Height())
		rad := s.MinRadius
		if s.MaxRadius > rad {
			rad = rand.Intn(s.MaxRadius-s.MinRadius) + s.MinRadius
		}
		s.disturb(grass, x, y, rad)
	}
}

func (s *SysDisturbance) disturb(grass *common.Grid[float32], x, y, rad int) {
	xmin := common.ClampInt(x-rad, 0, grass.Width())
	xmax := common.ClampInt(x+rad, 0, grass.Width())
	ymin := common.ClampInt(y-rad, 0, grass.Height())
	ymax := common.ClampInt(y+rad, 0, grass.Height())

	radSq := rad * rad

	for xx := xmin; xx < xmax; xx++ {
		dx := x - xx
		for yy := ymin; yy < ymax; yy++ {
			dy := y - yy
			if dx*dx+dy*dy <= radSq {
				grass.Set(xx, yy, s.TargetValue)
			}
		}
	}
}

// Finalize the system
func (s *SysDisturbance) Finalize(world *ecs.World) {}
