package box2d

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/mlange-42/arche-demo/common"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// DrawEntities is a system to draw Box2D bodies on an [Image] resource.
type DrawEntities struct {
	canvas generic.Resource[common.EbitenImage]
	images generic.Resource[Images]
	filter generic.Filter1[Body]
}

// InitializeUI the system
func (s *DrawEntities) InitializeUI(world *ecs.World) {
	s.canvas = generic.NewResource[common.EbitenImage](world)
	s.images = generic.NewResource[Images](world)
	s.filter = *generic.NewFilter1[Body]()
}

// UpdateUI the system
func (s *DrawEntities) UpdateUI(world *ecs.World) {
	grey := color.RGBA{160, 160, 160, 255}

	canvas := s.canvas.Get()
	img := canvas.Image

	query := s.filter.Query(world)
	for query.Next() {
		bodyComp := query.Get()
		pos := bodyComp.Body.GetPosition()
		r := bodyComp.Radius

		vector.DrawFilledCircle(img, float32(pos.X), float32(pos.Y), float32(r), grey, true)
	}
}

// PostUpdateUI the system
func (s *DrawEntities) PostUpdateUI(world *ecs.World) {}

// FinalizeUI the system
func (s *DrawEntities) FinalizeUI(world *ecs.World) {}
