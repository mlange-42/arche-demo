package boids

import (
	"math"

	"github.com/mlange-42/arche-demo/common"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// MoveEntities is a system that moves entities around.
type MoveEntities struct {
	ViewRadius float64

	canvas generic.Resource[common.Image]
	grid   generic.Resource[Grid]
	filter generic.Filter2[Position, Velocity]
	posMap generic.Map1[Position]

	neighbors []*GridEntry
}

// Initialize the system
func (s *MoveEntities) Initialize(world *ecs.World) {
	s.canvas = generic.NewResource[common.Image](world)
	s.grid = generic.NewResource[Grid](world)
	s.filter = *generic.NewFilter2[Position, Velocity]()
	s.posMap = generic.NewMap1[Position](world)

	s.neighbors = make([]*GridEntry, 0, 32)
}

// Update the system
func (s *MoveEntities) Update(world *ecs.World) {
	s.updateGrid(world)
	grid := s.grid.Get()

	canvas := s.canvas.Get()
	w := float64(canvas.Width)
	h := float64(canvas.Height)
	rad := s.ViewRadius

	query := s.filter.Query(world)
	cnt := 0
	for query.Next() {
		pos, vel := query.Get()

		s.processEntity(grid, query.Entity(), pos, vel, rad)

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
		cnt++
	}
}

func (s MoveEntities) updateGrid(world *ecs.World) {
	grid := s.grid.Get()

	grid.Clear()

	query := s.filter.Query(world)
	for query.Next() {
		pos, vel := query.Get()

		x, y := grid.ToCell(pos.X, pos.Y)
		grid.Add(x, y, query.Entity(), pos, vel)
	}
}

func (s MoveEntities) processEntity(grid *Grid, entity ecs.Entity, pos *Position, vel *Velocity, rad float64) {
	s.neighbors = s.neighbors[:0]

	x, y := pos.X, pos.Y
	r := int(math.Ceil(rad / float64(grid.CellSize)))
	radSq := rad * rad
	cx, cy := grid.ToCell(x, y)
	for i := -r; i <= r; i++ {
		xx := cx + i
		for j := -r; j <= r; j++ {
			yy := cy + j
			if !grid.Contains(xx, yy) {
				continue
			}
			boids := grid.Get(xx, yy)
			ln := len(boids)
			for k := 0; k < ln; k++ {
				entry := &boids[k]
				if entry.Entity != entity && common.DistanceSq(x, y, entry.X, entry.Y) <= radSq {
					s.neighbors = append(s.neighbors, entry)
				}
			}
		}
	}
}

// Finalize the system
func (s *MoveEntities) Finalize(world *ecs.World) {}
