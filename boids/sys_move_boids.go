package boids

import (
	"github.com/mlange-42/arche-demo/common"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// UISysDrawBoids is a system that draws ants.
type SysMoveBoids struct {
	Speed float64

	canvas generic.Resource[common.EbitenImage]
	filter generic.Filter2[Position, Heading]
}

// InitializeUI the system
func (s *SysMoveBoids) Initialize(world *ecs.World) {
	s.canvas = generic.NewResource[common.EbitenImage](world)
	s.filter = *generic.NewFilter2[Position, Heading]()
}

// UpdateUI the system
func (s *SysMoveBoids) Update(world *ecs.World) {
	//screen := s.canvas.Get()

	query := s.filter.Query(world)
	for query.Next() {
		pos, head := query.Get()
		dx, dy := head.Direction()
		pos.X += dx * s.Speed
		pos.Y += dy * s.Speed
	}
}

// FinalizeUI the system
func (s *SysMoveBoids) Finalize(world *ecs.World) {}
