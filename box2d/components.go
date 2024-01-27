package box2d

import (
	"github.com/ByteArena/box2d"
)

// Body component
type Body struct {
	Body   *box2d.B2Body
	Radius float64
}

// BoxWorld resource
type BoxWorld struct {
	World *box2d.B2World
}
