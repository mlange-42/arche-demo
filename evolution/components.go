package evolution

import (
	"image/color"
	"math"
	"math/rand"
)

// GeneNames for plotting.
var GeneNames = [7]string{
	"MaxAngle",
	"MinGrass",
	"Invest",
	"Offspring",
	"Red",
	"Green",
	"Blue",
}

// Position component
type Position struct {
	X float32
	Y float32
}

// Age component
type Age struct {
	TickOfBirth int64
}

// Energy component
type Energy struct {
	Energy float32
}

// Activity component
type Activity struct {
	IsGrazing bool
}

// Heading component
type Heading struct {
	Angle float32
}

// Direction returns the unit vector corresponding to the heading.
func (h *Heading) Direction() (float32, float32) {
	a := float64(h.Angle)
	return float32(math.Cos(a)), float32(math.Sin(a))
}

// Genotype components
type Genotype struct {
	Genes [7]float32
}

// Randomize all genes
func (g *Genotype) Randomize() {
	for i := range g.Genes {
		g.Genes[i] = rand.Float32()
	}
}

// Defaults sets all genes to default values
func (g *Genotype) Defaults() {
	for i := range g.Genes {
		g.Genes[i] = 0.5
	}
}

// Phenotype components
type Phenotype struct {
	MaxAngle  float32
	MinGrass  float32
	Invest    float32
	Offspring uint8
}

// From sets the Phenotype to a [Genotype].
func (p *Phenotype) From(g *Genotype) {
	p.MaxAngle = g.Genes[0] * 0.5 * math.Pi
	p.MinGrass = g.Genes[1] * 0.5
	p.Invest = g.Genes[2]
	p.Offspring = uint8(1 + g.Genes[3]*10)
}

// Color components
type Color struct {
	Color color.RGBA
}

// From sets the Colors to a [Genotype].
func (c *Color) From(g *Genotype) {
	c.Color.R = 50 + uint8(g.Genes[4]*200)
	c.Color.G = 50 + uint8(g.Genes[5]*200)
	c.Color.B = 50 + uint8(g.Genes[6]*200)
	c.Color.A = 255
}
