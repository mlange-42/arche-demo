package common

import (
	"math"
	"math/rand"
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

// SelectRoulette performs roulette selection
func SelectRoulette(probs []float64) (int, bool) {
	sum := 0.0
	for _, prob := range probs {
		sum += prob
	}
	if sum == 0 {
		return -1, false
	}
	r := rand.Float64() * sum

	cum := 0.0
	for i, prob := range probs {
		cum += prob
		if r <= cum {
			return i, true
		}
	}
	return -1, false
}
