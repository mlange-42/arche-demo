package boids

import (
	"math"
	"math/rand"

	"github.com/mlange-42/arche-demo/common"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// SysInitBoids system.
type SysInitBoids struct {
	Count int
}

// Initialize the system
func (s *SysInitBoids) Initialize(w *ecs.World) {
	screen := generic.NewResource[common.EbitenImage](w)
	scr := screen.Get()

	builder := generic.NewMap4[Position, Heading, Neighbors, Rand256](w)
	query := builder.NewBatchQ(s.Count)
	for query.Next() {
		pos, head, _, r256 := query.Get()
		pos.X, pos.Y = rand.Float64()*float64(scr.Width), rand.Float64()*float64(scr.Height)
		head.Angle = rand.Float64() * math.Pi * 2
		r256.R = uint8(rand.Int31n(256))
	}
}

// Update the system
func (s *SysInitBoids) Update(w *ecs.World) {}

// Finalize the system
func (s *SysInitBoids) Finalize(w *ecs.World) {}
