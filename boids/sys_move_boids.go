package boids

import (
	"github.com/mlange-42/arche-demo/common"
	"github.com/mlange-42/arche-model/resource"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// UISysDrawBoids is a system that draws ants.
type SysMoveBoids struct {
	Speed          float64
	UpdateInterval int

	SeparationDist   float64
	SeparationWeight float64
	CohesionWeight   float64
	AlignmentWeight  float64
	SpeedWeight      float64

	separationDistSq float64

	time   generic.Resource[resource.Tick]
	canvas generic.Resource[common.EbitenImage]
	filter generic.Filter4[Position, Velocity, Neighbors, Rand256]
}

// InitializeUI the system
func (s *SysMoveBoids) Initialize(world *ecs.World) {
	s.time = generic.NewResource[resource.Tick](world)
	s.canvas = generic.NewResource[common.EbitenImage](world)
	s.filter = *generic.NewFilter4[Position, Velocity, Neighbors, Rand256]()

	s.separationDistSq = s.SeparationDist * s.SeparationDist
}

// UpdateUI the system
func (s *SysMoveBoids) Update(world *ecs.World) {
	tick := s.time.Get().Tick
	modTick := uint8(tick % int64(s.UpdateInterval))

	query := s.filter.Query(world)
	for query.Next() {
		pos, vel, neigh, r256 := query.Get()

		if r256.R%uint8(s.UpdateInterval) == modTick {
			acc := s.separation(pos, neigh)
			acc.Add(s.cohesion(pos, neigh))
			acc.Add(s.alignment(vel, neigh))

			vel.Add(acc)

			ln := vel.Len()
			lnNew := (1-s.SpeedWeight)*ln + s.Speed*s.SpeedWeight
			vel.Norm(lnNew)
		}

		pos.X += vel.X
		pos.Y += vel.Y
	}
}

// FinalizeUI the system
func (s *SysMoveBoids) Finalize(world *ecs.World) {}

func (s *SysMoveBoids) separation(pos *Position, neigh *Neighbors) common.Vec2f {
	if len(neigh.Entities) == 0 {
		return common.Vec2f{}
	}

	dx, dy := 0.0, 0.0
	cnt := 0

	for _, n := range neigh.Entities {
		distSq := n.Vec2f.DistanceSq(pos.Vec2f)
		if distSq > s.separationDistSq {
			continue
		}
		x, y := pos.X-n.X, pos.Y-n.Y
		x, y, _ = common.Norm(x, y)
		dx += x
		dy += y
		cnt++
	}
	if cnt == 0 {
		return common.Vec2f{}
	}

	out := common.Vec2f{X: dx * s.SeparationWeight / float64(cnt), Y: dy * s.SeparationWeight / float64(cnt)}
	return out
}

func (s *SysMoveBoids) cohesion(pos *Position, neigh *Neighbors) common.Vec2f {
	return common.Vec2f{}
}

func (s *SysMoveBoids) alignment(vel *Velocity, neigh *Neighbors) common.Vec2f {
	return common.Vec2f{}
}
