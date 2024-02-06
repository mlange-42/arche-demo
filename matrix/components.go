package matrix

// Position component.
type Position struct {
	X int
	Y int
}

// Position component
type Letter struct {
	Letter rune
	Size   int
}

// Mover component
type Mover struct {
	Interval   uint16
	LastMove   int64
	PathLength int
}

// Fader component
type Fader struct {
	Start     int64
	Intensity float64
}
