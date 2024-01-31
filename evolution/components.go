package evolution

import (
	"image/color"
	"math"
	"math/rand"

	"github.com/mlange-42/arche-demo/common"
)

// GeneNames for plotting.
var GeneNames = [4]string{
	"MaxAngle",
	"MinGrass",
	"Invest",
	"Offspring",
}

// Position component
type Position struct {
	X float32
	Y float32
}

// Energy component
type Energy struct {
	Energy float32
}

// Grazing component
type Grazing struct{}

// Searching component
type Searching struct{}

// Heading component
type Heading struct {
	Angle float32
}

// Direction returns the unit vector oh the heading's angle.
func (h *Heading) Direction() (float32, float32) {
	a := float64(h.Angle)
	return float32(math.Cos(a)), float32(math.Sin(a))
}

// Genotype components
type Genotype struct {
	Genes [4]float32
}

// Randomize all genes
func (g *Genotype) Randomize() {
	g.Genes[0] = rand.Float32()
	g.Genes[1] = rand.Float32()
	g.Genes[2] = rand.Float32()
	g.Genes[3] = rand.Float32()
}

// Defaults sets all genes to default values
func (g *Genotype) Defaults() {
	g.Genes[0] = 0.5
	g.Genes[1] = 0.5
	g.Genes[2] = 0.5
	g.Genes[3] = 0.5
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

// Randomize all bands
func (c *Color) Randomize() {
	c.Color.R = common.RandBetweenUIn8(50, 250)
	c.Color.G = common.RandBetweenUIn8(50, 250)
	c.Color.B = common.RandBetweenUIn8(50, 250)
	c.Color.A = 255
}

// Defaults sets all bands to default values
func (c *Color) Defaults() {
	c.Color.R = 150
	c.Color.G = 150
	c.Color.B = 150
	c.Color.A = 255
}
