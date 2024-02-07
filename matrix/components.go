package matrix

// Position component.
type Position struct {
	X int
	Y int
}

// Position component.
type Letter struct {
	Letter rune
	Size   int
}

// Mover component for the bright letters that move downwards.
type Mover struct {
	Interval   uint16
	LastMove   int64
	PathLength int
}

// Fader component for the fading stationary letters.
type Fader struct {
	Intensity  float64
	NextChange int64
}
