package common

import (
	"fmt"
	"math"
)

// Vec2f is a 2d float vector
type Vec2f struct {
	X, Y float64
}

// Get returns the coordinate value for a specific dimension
func (v Vec2f) Get(dim int) float64 {
	switch dim {
	case 0:
		return v.X
	case 1:
		return v.Y
	default:
		panic(fmt.Sprintf("Invalid dimension %d", dim))
	}
}

// Distance to another vector
func (v Vec2f) Distance(other Vec2f) float64 {
	return math.Hypot(v.X-other.X, v.Y-other.Y)
}

// DistanceSq to another vector
func (v Vec2f) DistanceSq(other Vec2f) float64 {
	dx := v.X - other.X
	dy := v.Y - other.Y
	return dx*dx + dy*dy
}

// Add adds a vector in-place
func (v *Vec2f) Add(other Vec2f) {
	v.X += other.X
	v.Y += other.Y
}

// Sub subtracts a vector in-place
func (v *Vec2f) Sub(other Vec2f) {
	v.X -= other.X
	v.Y -= other.Y
}

// Div divides a vector in-place
func (v *Vec2f) Div(factor float64) {
	v.X /= factor
	v.Y /= factor
}

// Mul multiplies a vector in-place
func (v *Vec2f) Mul(factor float64) {
	v.X *= factor
	v.Y *= factor
}

// IsZero adds a vector in-place
func (v *Vec2f) IsZero() bool {
	return v.X == 0 && v.Y == 0
}

// Norm normalizes a vector to the given length and returns the original length.
func (v *Vec2f) Norm(length float64) float64 {
	l := math.Hypot(v.X, v.Y)
	f := length / l
	v.X *= f
	v.Y *= f
	return l
}

// LenSq returns the squared length of the vector
func (v Vec2f) LenSq() float64 {
	return v.X*v.X + v.Y*v.Y
}

// Len returns the length of the vector
func (v Vec2f) Len() float64 {
	return math.Hypot(v.X, v.Y)
}

// Len returns the length of the vector
func (v Vec2f) Angle() float64 {
	return math.Atan2(v.Y, v.X)
}

// Copy copies the vector
func (v *Vec2f) Copy() Vec2f {
	return Vec2f{v.X, v.Y}
}
