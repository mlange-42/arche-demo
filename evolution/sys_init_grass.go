package evolution

import (
	"time"

	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
	"github.com/ojrac/opensimplex-go"
)

// SysInitGrass is a system to initialize the [Grass] resource.
type SysInitGrass struct {
	Count int
}

// Initialize the system
func (s *SysInitGrass) Initialize(world *ecs.World) {
	grassRes := generic.NewResource[Grass](world)
	grass := grassRes.Get()

	w, h := grass.Grass.Width(), grass.Grass.Height()
	noise := opensimplex.NewNormalized32(int64(time.Now().Nanosecond()))

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			grass.Grass.Set(x, y, noise.Eval2(float32(x), float32(y)))
		}
	}
}

// Update the system
func (s *SysInitGrass) Update(world *ecs.World) {}

// Finalize the system
func (s *SysInitGrass) Finalize(world *ecs.World) {}
