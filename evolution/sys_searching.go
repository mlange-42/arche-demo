package evolution

import (
	"math"
	"math/rand"

	"github.com/mlange-42/arche-demo/common"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// SysSearching is a system to move around searching entities.
type SysSearching struct {
	MaxSpeed float32

	grass  generic.Resource[Grass]
	filter generic.Filter4[Position, Activity, Heading, Phenotype]
}

// Initialize the system
func (s *SysSearching) Initialize(world *ecs.World) {
	s.grass = generic.NewResource[Grass](world)
	s.filter = *generic.NewFilter4[Position, Activity, Heading, Phenotype]()
}

// Update the system
func (s *SysSearching) Update(world *ecs.World) {
	grass := &s.grass.Get().Grass

	query := s.filter.Query(world)
	for query.Next() {
		pos, act, head, pt := query.Get()
		if act.IsGrazing {
			continue
		}

		head.Angle += rand.Float32()*pt.MaxAngle*2 - pt.MaxAngle
		head.Angle = common.NormAngle32(head.Angle)

		dx, dy := head.Direction()
		pos.X += dx * s.MaxSpeed
		pos.Y += dy * s.MaxSpeed

		if !grass.IsInBounds(float64(pos.X), float64(pos.Y)) {
			pos.X -= dx * s.MaxSpeed
			pos.Y -= dy * s.MaxSpeed
			head.Angle = common.NormAngle32(head.Angle + math.Pi)
		}
	}
}

// Finalize the system
func (s *SysSearching) Finalize(world *ecs.World) {}
