package box2d

import (
	"math"

	"github.com/ByteArena/box2d"
	"github.com/mlange-42/arche-demo/common"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// SysPhysics is a system that applies an impulse to all Box2D bodies
// that are within a certain distance to the mouse.
type SysPhysics struct {
	MinFleeDistance float64
	MaxFleeDistance float64
	ForceScale      float64
	mouse           generic.Resource[common.Mouse]
	filter          generic.Filter1[Body]
}

// Initialize the system
func (s *SysPhysics) Initialize(world *ecs.World) {
	s.mouse = generic.NewResource[common.Mouse](world)
	s.filter = *generic.NewFilter1[Body]()
}

// Update the system
func (s *SysPhysics) Update(world *ecs.World) {
	mouse := s.mouse.Get()
	if !mouse.IsInside {
		return
	}

	minDist := s.MinFleeDistance
	distRange := s.MaxFleeDistance - minDist

	query := s.filter.Query(world)
	for query.Next() {
		body := query.Get()
		pos := body.Body.GetPosition()
		repX, repY, repDist := common.Norm(pos.X-float64(mouse.X), pos.Y-float64(mouse.Y))
		repFac := math.Min(1.0-((repDist-minDist)/distRange), 1.0) * body.Body.M_mass
		if repFac > 0 {
			body.Body.ApplyLinearImpulseToCenter(
				box2d.B2Vec2{X: repX * repFac * s.ForceScale, Y: repY * repFac * s.ForceScale},
				true,
			)
		}
	}
}

// Finalize the system
func (s *SysPhysics) Finalize(world *ecs.World) {}
