package boids

import (
	"github.com/mlange-42/arche-demo/common"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// MoveEntities is a system that moves entities around.
type MoveEntities struct {
	canvas generic.Resource[common.Image]
	grid   generic.Resource[Grid]
	filter generic.Filter2[Position, Velocity]
	posMap generic.Map1[Position]
	relMap generic.Map[CurrentCell]

	toMove []ecs.Entity
}

// Initialize the system
func (s *MoveEntities) Initialize(world *ecs.World) {
	s.canvas = generic.NewResource[common.Image](world)
	s.grid = generic.NewResource[Grid](world)
	s.filter = *generic.NewFilter2[Position, Velocity]()
	s.posMap = generic.NewMap1[Position](world)
	s.relMap = generic.NewMap[CurrentCell](world)

	s.toMove = make([]ecs.Entity, 0, 64)
}

// Update the system
func (s *MoveEntities) Update(world *ecs.World) {
	grid := s.grid.Get()
	canvas := s.canvas.Get()
	w := float64(canvas.Width)
	h := float64(canvas.Height)

	query := s.filter.Query(world)
	for query.Next() {
		pos, vel := query.Get()
		cxOld, cyOld := grid.ToCell(pos.X, pos.Y)

		pos.X += vel.X
		pos.Y += vel.Y

		if pos.X < 0 || pos.X >= w {
			vel.X *= -1
			pos.X += vel.X * 2
		}
		if pos.Y < 0 || pos.Y >= h {
			vel.Y *= -1
			pos.Y += vel.Y * 2
		}

		cx, cy := grid.ToCell(pos.X, pos.Y)
		if cx != cxOld || cy != cyOld {
			s.toMove = append(s.toMove, query.Entity())
		}
	}

	for _, e := range s.toMove {
		pos := s.posMap.Get(e)
		cx, cy := grid.ToCell(pos.X, pos.Y)
		cell := grid.Cells[cx][cy]
		s.relMap.SetRelation(e, cell)
	}

	s.toMove = s.toMove[:0]
}

// Finalize the system
func (s *MoveEntities) Finalize(world *ecs.World) {}
