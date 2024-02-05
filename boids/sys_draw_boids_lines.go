package boids

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/mlange-42/arche-demo/common"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// UISysDrawBoidsLines is a system that draws boids.
type UISysDrawBoidsLines struct {
	canvas generic.Resource[common.EbitenImage]
	images generic.Resource[Images]
	filter generic.Filter2[Position, Velocity]
}

// InitializeUI the system
func (s *UISysDrawBoidsLines) InitializeUI(world *ecs.World) {
	s.canvas = generic.NewResource[common.EbitenImage](world)
	s.images = generic.NewResource[Images](world)
	s.filter = *generic.NewFilter2[Position, Velocity]()
}

// UpdateUI the system
func (s *UISysDrawBoidsLines) UpdateUI(world *ecs.World) {
	//images := s.images.Get()
	canvas := s.canvas.Get()
	img := canvas.Image

	col := color.RGBA{R: 0, G: 255, B: 0, A: 255}

	img.Clear()

	query := s.filter.Query(world)
	for query.Next() {
		pos, vel := query.Get()
		dx, dy, _ := common.Norm(vel.X, vel.Y)
		vector.StrokeLine(img, float32(pos.X), float32(pos.Y), float32(pos.X+3*dx), float32(pos.Y+3*dy), 1, col, false)
	}
}

// PostUpdateUI the system
func (s *UISysDrawBoidsLines) PostUpdateUI(world *ecs.World) {}

// FinalizeUI the system
func (s *UISysDrawBoidsLines) FinalizeUI(world *ecs.World) {}
