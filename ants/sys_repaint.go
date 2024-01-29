package ants

import (
	"github.com/mlange-42/arche-demo/common"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// SysRepaint is a simple system that paints an [Image] resource to a [common.Canvas].
type SysRepaint struct {
	canvas generic.Resource[common.Image]
	screen generic.Resource[common.EbitenImage]
}

// InitializeUI the system
func (s *SysRepaint) InitializeUI(world *ecs.World) {
	s.canvas = generic.NewResource[common.Image](world)
	s.screen = generic.NewResource[common.EbitenImage](world)
}

// UpdateUI the system
func (s *SysRepaint) UpdateUI(world *ecs.World) {
	canvas := s.canvas.Get()
	screen := s.screen.Get()
	screen.Image.WritePixels(canvas.Image.Pix)
}

// PostUpdateUI the system
func (s *SysRepaint) PostUpdateUI(world *ecs.World) {}

// FinalizeUI the system
func (s *SysRepaint) FinalizeUI(world *ecs.World) {}
