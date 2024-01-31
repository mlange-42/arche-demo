package common

import (
	"math"
	"math/rand"
)

// DegToRad conversion factor.
const DegToRad = math.Pi / 180.0

// RadToDeg conversion factor.
const RadToDeg = 180.0 / math.Pi

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

// NormAngle brings an angle into range [0, 2*PI).
func NormAngle(angle float64) float64 {
	if angle < 0 {
		return math.Pi*2 - math.Mod(angle, math.Pi*2)
	}
	return math.Mod(angle, math.Pi*2)
}

// NormAngle32 brings an angle into range [0, 2*PI).
func NormAngle32(angle float32) float32 {
	if angle < 0 {
		return float32(math.Pi*2 + math.Mod(float64(angle), math.Pi*2))
	}
	return float32(math.Mod(float64(angle), math.Pi*2))
}

// ClampInt clamps to [low, high].
func ClampInt(v, low, high int) int {
	if v < low {
		return low
	}
	if v > high {
		return high
	}
	return v
}

// Clamp clamps to [low, high].
func Clamp(v, low, high float64) float64 {
	if v < low {
		return low
	}
	if v > high {
		return high
	}
	return v
}

// Clamp clamps to [low, high].
func Clamp32(v, low, high float32) float32 {
	if v < low {
		return low
	}
	if v > high {
		return high
	}
	return v
}
