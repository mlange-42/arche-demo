package matrix

// Position component.
type Position struct {
	X int
	Y int
}

// Position component.
type Letter struct {
	Letter rune
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

// Message component.
type Message struct {
	Message int
	Index   int
}

// ForcedLetter component.
type ForcedLetter struct {
	Letter    rune
	Active    bool
	Traversed bool
}

type LetterForcer struct {
	Message  int
	TickDone int64
}
