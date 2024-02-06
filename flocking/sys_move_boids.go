package flocking

import (
	"math"

	"github.com/mlange-42/arche-demo/common"
	"github.com/mlange-42/arche-model/resource"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// UISysDrawBoids is a system that draws ants.
type SysMoveBoids struct {
	Speed          float64
	MaxAcc         float64
	UpdateInterval int

	SeparationDist  float64
	SeparationAngle float64
	CohesionAngle   float64
	AlignmentAngle  float64

	WallDist  float64
	WallAngle float64

	MouseDist  float64
	MouseAngle float64

	separationDistSq float64
	mouseDistSq      float64
	separationAngle  float64
	cohesionAngle    float64
	alignmentAngle   float64
	wallAngle        float64
	mouseAngle       float64

	time   generic.Resource[resource.Tick]
	mouse  generic.Resource[common.Mouse]
	canvas generic.Resource[common.EbitenImage]
	filter generic.Filter4[Position, Heading, Neighbors, Rand256]
}

// InitializeUI the system
func (s *SysMoveBoids) Initialize(world *ecs.World) {
	s.time = generic.NewResource[resource.Tick](world)
	s.mouse = generic.NewResource[common.Mouse](world)
	s.canvas = generic.NewResource[common.EbitenImage](world)
	s.filter = *generic.NewFilter4[Position, Heading, Neighbors, Rand256]()

	s.separationDistSq = s.SeparationDist * s.SeparationDist
	s.mouseDistSq = s.MouseDist * s.MouseDist

	s.separationAngle = s.SeparationAngle * common.DegToRad
	s.cohesionAngle = s.CohesionAngle * common.DegToRad
	s.alignmentAngle = s.AlignmentAngle * common.DegToRad
	s.wallAngle = s.WallAngle * common.DegToRad
	s.mouseAngle = s.MouseAngle * common.DegToRad
}

// UpdateUI the system
func (s *SysMoveBoids) Update(world *ecs.World) {
	tick := s.time.Get().Tick
	modTick := uint8(tick % int64(s.UpdateInterval))

	screen := s.canvas.Get()
	mouse := s.mouse.Get()

	query := s.filter.Query(world)
	for query.Next() {
		pos, head, neigh, r256 := query.Get()
		_ = neigh

		if r256.R%uint8(s.UpdateInterval) == modTick {
			turn := 0.0
			turn += common.Clamp(s.separation(*pos, head.Angle, neigh), -s.separationAngle, s.separationAngle)
			turn += common.Clamp(s.cohesion(*pos, head.Angle, neigh), -s.cohesionAngle, s.cohesionAngle)
			turn += common.Clamp(s.alignment(head.Angle, neigh), -s.separationAngle, s.separationAngle)

			av, avLen := s.wallAvoidance(*pos, head.Angle, float64(screen.Width), float64(screen.Height))
			turn += common.Clamp(av, -s.wallAngle*avLen, s.wallAngle*avLen)

			av, avLen = s.mouseAvoidance(*pos, head.Angle, mouse)
			turn += common.Clamp(av, -s.wallAngle*avLen, s.wallAngle*avLen)

			head.Angle += turn
		}
		v := head.Direction()
		pos.X += v.X * s.Speed
		pos.Y += v.Y * s.Speed
	}
}

// FinalizeUI the system
func (s *SysMoveBoids) Finalize(world *ecs.World) {}

func (s *SysMoveBoids) separation(pos Position, angle float64, neigh *Neighbors) float64 {
	if neigh.Count == 0 {
		return 0
	}

	n := &neigh.Entities[0]
	distSq := n.Vec2f.DistanceSq(pos.Vec2f)
	if distSq > s.separationDistSq {
		return 0
	}
	ang := math.Atan2(n.Y-pos.Y, n.X-pos.X)
	return -common.SubtractHeadings(ang, angle)
}

func (s *SysMoveBoids) cohesion(pos Position, angle float64, neigh *Neighbors) float64 {
	if neigh.Count == 0 {
		return 0
	}

	cx, cy := 0.0, 0.0

	for i := 0; i < neigh.Count; i++ {
		n := &neigh.Entities[i]
		cx += n.X
		cy += n.Y
	}
	cx /= float64(neigh.Count)
	cy /= float64(neigh.Count)
	ang := math.Atan2(cy-pos.Y, cx-pos.X)
	return common.SubtractHeadings(ang, angle)
}

func (s *SysMoveBoids) alignment(angle float64, neigh *Neighbors) float64 {
	if neigh.Count == 0 {
		return 0
	}
	dx, dy := 0.0, 0.0
	for i := 0; i < neigh.Count; i++ {
		n := &neigh.Entities[i]
		dx += math.Cos(n.Heading)
		dy += math.Sin(n.Heading)
	}
	dx /= float64(neigh.Count)
	dy /= float64(neigh.Count)

	ang := math.Atan2(dy, dx)
	return common.SubtractHeadings(ang, angle)
}

func (s *SysMoveBoids) wallAvoidance(pos Position, angle float64, w, h float64) (float64, float64) {
	target := common.Vec2f{}
	if pos.X < s.WallDist {
		target.X += (s.WallDist - pos.X) / s.WallDist
	}
	if pos.Y < s.WallDist {
		target.Y += (s.WallDist - pos.Y) / s.WallDist
	}
	if pos.X > w-s.WallDist {
		target.X -= (s.WallDist - (w - pos.X)) / s.WallDist
	}
	if pos.Y > h-s.WallDist {
		target.Y -= (s.WallDist - (h - pos.Y)) / s.WallDist
	}
	if target.X == 0 && target.Y == 0 {
		return 0, 0
	}

	ang := target.Angle()
	return common.SubtractHeadings(ang, angle), target.Len()
}

func (s *SysMoveBoids) mouseAvoidance(pos Position, angle float64, mouse *common.Mouse) (float64, float64) {
	if !mouse.IsInside {
		return 0, 0
	}
	mx, my := float64(mouse.X), float64(mouse.Y)
	distSq := pos.DistanceSq(common.Vec2f{X: mx, Y: my})
	if distSq > s.mouseDistSq {
		return 0, 0
	}

	ang := math.Atan2(my-pos.Y, mx-pos.X)
	return -common.SubtractHeadings(ang, angle), 1 - math.Sqrt(distSq)/s.MouseDist
}
