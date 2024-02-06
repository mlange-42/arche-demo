package boids

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/mlange-42/arche-demo/common"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// UISysDrawBoids is a system that draws a single boid and its neighbors.
type UISysDrawBoid struct {
	Radius float32

	canvas generic.Resource[common.EbitenImage]
	filter generic.Filter2[Position, Neighbors]
}

// InitializeUI the system
func (s *UISysDrawBoid) InitializeUI(world *ecs.World) {
	s.canvas = generic.NewResource[common.EbitenImage](world)
	s.filter = *generic.NewFilter2[Position, Neighbors]()
}

// UpdateUI the system
func (s *UISysDrawBoid) UpdateUI(world *ecs.World) {
	canvas := s.canvas.Get()
	img := canvas.Image

	col := color.RGBA{R: 160, G: 160, B: 160, A: 255}

	query := s.filter.Query(world)
	if query.Next() {
		pos, neigh := query.Get()

		vector.StrokeCircle(img, float32(pos.X), float32(pos.Y), s.Radius, 1, col, true)
		for i := 0; i < neigh.Count; i++ {
			n := &neigh.Entities[i]
			vector.StrokeLine(img, float32(pos.X), float32(pos.Y), float32(n.X), float32(n.Y), 1, col, true)
		}
	}
	query.Close()
}

// PostUpdateUI the system
func (s *UISysDrawBoid) PostUpdateUI(world *ecs.World) {}

// FinalizeUI the system
func (s *UISysDrawBoid) FinalizeUI(world *ecs.World) {}
