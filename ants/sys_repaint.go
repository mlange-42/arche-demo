package ants

import (
	"github.com/mlange-42/arche-demo/common"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// UISysRepaint is a simple system that paints an [Image] resource to a [common.Canvas].
type UISysRepaint struct {
	canvas generic.Resource[common.Image]
	screen generic.Resource[common.EbitenImage]
}

// InitializeUI the system
func (s *UISysRepaint) InitializeUI(world *ecs.World) {
	s.canvas = generic.NewResource[common.Image](world)
	s.screen = generic.NewResource[common.EbitenImage](world)
}

// UpdateUI the system
func (s *UISysRepaint) UpdateUI(world *ecs.World) {
	canvas := s.canvas.Get()
	screen := s.screen.Get()
	screen.Image.WritePixels(canvas.Image.Pix)
}

// PostUpdateUI the system
func (s *UISysRepaint) PostUpdateUI(world *ecs.World) {}

// FinalizeUI the system
func (s *UISysRepaint) FinalizeUI(world *ecs.World) {}
