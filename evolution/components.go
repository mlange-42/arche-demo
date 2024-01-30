package evolution

// Position component
type Position struct {
	X float32
	Y float32
}

// Heading component
type Heading struct {
	Angle float32
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
	R uint8
	G uint8
	B uint8
}
