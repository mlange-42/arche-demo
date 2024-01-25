package main

import (
	"math"
	"math/rand"

	"github.com/mlange-42/arche-model/resource"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// Scouting system
type Scouting struct {
	MaxRotation  float64
	MaxScoutTime int64
	canvas       generic.Resource[Image]
	patches      generic.Resource[Patches]
	params       generic.Resource[Params]
	time         generic.Resource[resource.Tick]
	filter       generic.Filter4[Position, Direction, Scout, Random256]
}

// Initialize the system
func (s *Scouting) Initialize(world *ecs.World) {
	s.canvas = generic.NewResource[Image](world)
	s.patches = generic.NewResource[Patches](world)
	s.params = generic.NewResource[Params](world)
	s.time = generic.NewResource[resource.Tick](world)
	s.filter = *generic.NewFilter4[Position, Direction, Scout, Random256]()
}

// Update the system
func (s *Scouting) Update(world *ecs.World) {
	canvas := s.canvas.Get()
	patches := s.patches.Get()

	w := float64(canvas.Width)
	h := float64(canvas.Height)

	maxSpeed := s.params.Get().MaxBeeSpeed
	maxAng := (s.MaxRotation * math.Pi / 180.0) / 2
	tick := s.time.Get().Tick

	query := s.filter.Query(world)
	for query.Next() {
		pos, dir, scout, r256 := query.Get()

		if tick%4 == int64(r256.Value)%4 {
			dir.X, dir.Y = rotate(dir.X, dir.Y, rand.Float64()*2*maxAng-maxAng)

			cx, cy := patches.ToCell(pos.X, pos.Y)
			patch := patches.Patches[cx][cy]
			if !patch.IsZero() {
				// forage
			}
			if scout.Start+s.MaxScoutTime > tick {
				// return to hive
			}
		}

		pos.X += dir.X * maxSpeed
		pos.Y += dir.Y * maxSpeed

		if pos.X < 0 || pos.X >= w {
			dir.X *= -1
			pos.X += dir.X * maxSpeed * 2
		}
		if pos.Y < 0 || pos.Y >= h {
			dir.Y *= -1
			pos.Y += dir.Y * maxSpeed * 2
		}
	}
}

// Finalize the system
func (s *Scouting) Finalize(world *ecs.World) {}
