package logo

import (
	"math"
	"math/rand"

	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// MoveEntities is a system that moves entities around.
// Entities accelerate towards their [Target] position,
// as well as away from the mouse if it is within a certain distance.
type MoveEntities struct {
	// Maximum speed of an entity.
	MaxSpeed float64
	// Maximum acceleration of an entity.
	MaxAcc float64
	// Maximum acceleration of an entity when fleeing.
	MaxAccFlee float64
	// Minimum flee distance, with maximum fleeing acceleration.
	MinFleeDistance float64
	// Maximum fleeing distance. Beyond that distance, entities don't flee.
	MaxFleeDistance float64
	// Dampening of entity movement.
	Damp    float64
	mouse   generic.Resource[MouseListener]
	filter  generic.Filter3[Position, Velocity, Target]
	counter uint64
}

// Initialize the system
func (s *MoveEntities) Initialize(world *ecs.World) {
	s.mouse = generic.NewResource[MouseListener](world)
	s.filter = *generic.NewFilter3[Position, Velocity, Target]()
}

// Update the system
func (s *MoveEntities) Update(world *ecs.World) {
	listener := s.mouse.Get()
	mouse := listener.Mouse
	mouseInside := listener.MouseInside

	minDist := s.MinFleeDistance
	distRange := s.MaxFleeDistance - minDist

	query := s.filter.Query(world)
	for query.Next() {
		pos, vel, trg := query.Get()

		attrX, attrY, _ := s.norm(trg.X-pos.X, trg.Y-pos.Y)

		vel.X += attrX * s.MaxAcc
		vel.Y += attrY * s.MaxAcc

		if mouseInside {
			repX, repY, repDist := s.norm(pos.X-mouse.X, pos.Y-mouse.Y)
			repFac := math.Min(1.0-((repDist-minDist)/distRange), 1.0)
			if repFac > 0 {
				vel.X += repX * s.MaxAccFlee * repFac
				vel.Y += repY * s.MaxAccFlee * repFac
			}
		}

		velAbs := vel.X*vel.X + vel.Y*vel.Y
		if velAbs > 1.0 {
			velAbs := math.Sqrt(velAbs)
			vel.X /= velAbs
			vel.Y /= velAbs
			velAbs = 1.0
		}
		if s.counter%23 == 0 {
			vel.X += rand.NormFloat64() * velAbs * 0.2
			vel.Y += rand.NormFloat64() * velAbs * 0.2
		}

		vel.X *= s.Damp
		vel.Y *= s.Damp

		pos.X += vel.X * s.MaxSpeed
		pos.Y += vel.Y * s.MaxSpeed

		s.counter++
	}
}

func (s *MoveEntities) norm(dx, dy float64) (float64, float64, float64) {
	len := math.Sqrt(dx*dx + dy*dy)
	if len == 0 {
		return 0, 0, 0
	}
	return dx / len, dy / len, len
}

// Finalize the system
func (s *MoveEntities) Finalize(world *ecs.World) {}
