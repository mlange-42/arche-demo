package boids

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mlange-42/arche-demo/common"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// UISysDrawBoids is a system that draws boids.
type UISysDrawBoids struct {
	canvas generic.Resource[common.EbitenImage]
	images generic.Resource[Images]
	filter generic.Filter2[Position, Heading]
}

// InitializeUI the system
func (s *UISysDrawBoids) InitializeUI(world *ecs.World) {
	s.canvas = generic.NewResource[common.EbitenImage](world)
	s.images = generic.NewResource[Images](world)
	s.filter = *generic.NewFilter2[Position, Heading]()
}

// UpdateUI the system
func (s *UISysDrawBoids) UpdateUI(world *ecs.World) {
	images := s.images.Get()
	canvas := s.canvas.Get()
	img := canvas.Image

	xOff, yOff := float64(images.Boid.Bounds().Dx())/2, float64(images.Boid.Bounds().Dy())/2

	img.Clear()

	opts := ebiten.DrawImageOptions{
		GeoM:   ebiten.GeoM{},
		Filter: ebiten.FilterLinear,
	}

	query := s.filter.Query(world)
	for query.Next() {
		pos, head := query.Get()

		opts.GeoM.Reset()
		opts.GeoM.Translate(-xOff, -yOff)
		opts.GeoM.Rotate(head.Angle)
		opts.GeoM.Translate(pos.X, pos.Y)
		img.DrawImage(images.Boid, &opts)
	}
}

// PostUpdateUI the system
func (s *UISysDrawBoids) PostUpdateUI(world *ecs.World) {}

// FinalizeUI the system
func (s *UISysDrawBoids) FinalizeUI(world *ecs.World) {}
