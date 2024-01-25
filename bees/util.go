package main

import "math"

func norm(dx, dy float64) (float64, float64, float64) {
	len := math.Sqrt(dx*dx + dy*dy)
	if len == 0 {
		return 0, 0, 0
	}
	return dx / len, dy / len, len
}

func rotate(x, y, angle float64) (float64, float64) {
	sin := math.Sin(angle)
	cos := math.Cos(angle)
	return cos*x - sin*y, sin*x + cos*y
}
