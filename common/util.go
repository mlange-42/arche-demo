package common

import (
	"math"
	"syscall/js"
)

// Norm normalizes a vector. The third return value is the original vector's length.
func Norm(dx, dy float64) (float64, float64, float64) {
	len := math.Sqrt(dx*dx + dy*dy)
	if len == 0 {
		return 0, 0, 0
	}
	return dx / len, dy / len, len
}

// Rotate rotates a vector.
func Rotate(x, y, angle float64) (float64, float64) {
	sin := math.Sin(angle)
	cos := math.Cos(angle)
	return cos*x - sin*y, sin*x + cos*y
}

// DistanceSq calculates the squared distance between two points.
func DistanceSq(x1, y1, x2, y2 float64) float64 {
	dx := x1 - x2
	dy := y1 - y2
	return dx*dx + dy*dy
}

// Distance calculates the distance between two points.
func Distance(x1, y1, x2, y2 float64) float64 {
	dx := x1 - x2
	dy := y1 - y2
	return math.Sqrt(dx*dx + dy*dy)
}

// RemoveElementByID removes an HTML element by ID.
func removeElementByID(doc js.Value, id string) {
	elem := doc.Call("getElementById", id)
	elem.Call("remove")
}
