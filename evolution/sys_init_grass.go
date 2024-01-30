package evolution

import (
	"math"
	"time"

	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
	"github.com/ojrac/opensimplex-go"
)

// SysInitGrass is a system to initialize the [Grass] resource.
type SysInitGrass struct {
	Frequency float32
	Octaves   int
	Cutoff    float32
}

// Initialize the system
func (s *SysInitGrass) Initialize(world *ecs.World) {
	grassRes := generic.NewResource[Grass](world)
	grass := grassRes.Get()

	w, h := grass.Grass.Width(), grass.Grass.Height()
	noise := opensimplex.NewNormalized32(int64(time.Now().Nanosecond()))

	var max float32 = 0.0
	for o := 0; o < s.Octaves; o++ {
		max += float32(math.Pow(0.5, float64(o)))
	}

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			var v float32 = 0.0
			for o := 0; o < s.Octaves; o++ {
				fac := float32(math.Pow(0.5, float64(o)))
				freq := s.Frequency / fac
				v += noise.Eval2(float32(x)*freq, float32(y)*freq) * fac
			}
			v /= max

			v = (v - s.Cutoff) / (1.0 - s.Cutoff)
			if v < 0 {
				v = 0
			}

			grass.Grass.Set(x, y, v)
			grass.Growth.Set(x, y, v)
		}
	}
}

// Update the system
func (s *SysInitGrass) Update(world *ecs.World) {}

// Finalize the system
func (s *SysInitGrass) Finalize(world *ecs.World) {}
