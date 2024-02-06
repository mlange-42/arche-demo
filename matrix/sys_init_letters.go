package matrix

import (
	"math/rand"

	"github.com/mlange-42/arche-demo/common"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// SysInitLetters is a system to initialize entities.
type SysInitLetters struct {
	Count int
}

// Initialize the system
func (s *SysInitLetters) Initialize(world *ecs.World) {
	canvas := generic.NewResource[common.EbitenImage](world)
	screen := canvas.Get()
	builder := generic.NewMap2[Position, Letter](world)

	letters := []rune(characters)
	query := builder.NewBatchQ(s.Count)
	for query.Next() {
		pos, let := query.Get()

		pos.X = rand.Float64()*float64(screen.Width-20) + 10
		pos.Y = rand.Float64()*float64(screen.Height-20) + 10
		let.Letter = letters[rand.Intn(len(letters))]
		let.Size = rand.Intn(len(fontSizes))
	}
}

// Update the system
func (s *SysInitLetters) Update(world *ecs.World) {}

// Finalize the system
func (s *SysInitLetters) Finalize(world *ecs.World) {}
