package evolution

import (
	"image/color"
	"math"
)

// Position component
type Position struct {
	X float32
	Y float32
}

// Energy component
type Energy struct {
	Energy float32
}

// Grazing component
type Grazing struct{}

// Searching component
type Searching struct{}

// Heading component
type Heading struct {
	Angle float32
}

// Direction returns the unit vector oh the heading's angle.
func (h *Heading) Direction() (float32, float32) {
	a := float64(h.Angle)
	return float32(math.Cos(a)), float32(math.Sin(a))
}

// Genes components
type Genes struct {
	MaxAngle  float32
	MinGrass  float32
	Invest    float32
	Offspring uint8
}

// Color components
type Color struct {
	Color color.RGBA
}
