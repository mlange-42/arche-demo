package main

import (
	"math/rand"

	"github.com/ByteArena/box2d"
	"github.com/mlange-42/arche-demo/common"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// InitEntities system
type InitEntities struct {
	Count       int
	Restitution float64
}

// Initialize the system
func (s *InitEntities) Initialize(world *ecs.World) {
	worldRes := generic.NewResource[BoxWorld](world)
	w := worldRes.Get().World
	canvasRes := generic.NewResource[common.Canvas](world)
	canvas := canvasRes.Get()

	bd := box2d.MakeB2BodyDef()
	ground := w.CreateBody(&bd)

	vs := make([]box2d.B2Vec2, 4)
	vs[0].Set(0.0, 0.0)
	vs[1].Set(float64(canvas.Width), 0.0)
	vs[2].Set(float64(canvas.Width), float64(canvas.Height))
	vs[3].Set(0, float64(canvas.Height))
	shape := box2d.MakeB2ChainShape()
	shape.CreateLoop(vs, 4)
	ground.CreateFixture(&shape, 0.0)

	builder := generic.NewMap1[Body](world)
	query := builder.NewBatchQ(s.Count)

	for query.Next() {
		bodyComp := query.Get()

		bd := box2d.MakeB2BodyDef()
		bd.Position.Set(
			rand.Float64()*float64(canvas.Width)*0.8+float64(canvas.Width)*0.1,
			rand.Float64()*float64(canvas.Height)*0.8+float64(canvas.Height)*0.1,
		)
		bd.Type = box2d.B2BodyType.B2_dynamicBody
		bd.AllowSleep = false

		body := w.CreateBody(&bd)
		body.SetLinearVelocity(box2d.B2Vec2{X: rand.NormFloat64() * 100, Y: rand.NormFloat64() * 100})

		shape := box2d.MakeB2CircleShape()
		shape.M_radius = rand.Float64()*10 + 5

		fd := box2d.MakeB2FixtureDef()
		fd.Shape = &shape
		fd.Density = 20.0
		fd.Friction = 1.0
		fd.Restitution = s.Restitution
		body.CreateFixtureFromDef(&fd)

		bodyComp.Body = body
		bodyComp.Radius = shape.M_radius
	}
}

// Update the system
func (s *InitEntities) Update(world *ecs.World) {}

// Finalize the system
func (s *InitEntities) Finalize(world *ecs.World) {}
