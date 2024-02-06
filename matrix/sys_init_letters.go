package matrix

import (
	"math/rand"

	"github.com/mlange-42/arche-demo/common"
	"github.com/mlange-42/arche-model/resource"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// SysInitLetters is a system to initialize entities.
type SysInitLetters struct {
	SpawnProb       float64
	MinMoveInterval int
	MaxMoveInterval int
	MinGap          int

	time    generic.Resource[resource.Tick]
	grid    generic.Resource[LetterGrid]
	canvas  generic.Resource[common.EbitenImage]
	letters generic.Resource[Letters]
	builder generic.Map3[Position, Letter, Mover]

	releases []int64
}

// Initialize the system
func (s *SysInitLetters) Initialize(world *ecs.World) {
	s.time = generic.NewResource[resource.Tick](world)
	s.grid = generic.NewResource[LetterGrid](world)
	s.canvas = generic.NewResource[common.EbitenImage](world)
	s.letters = generic.NewResource[Letters](world)
	s.builder = generic.NewMap3[Position, Letter, Mover](world)

	s.releases = make([]int64, s.grid.Get().Faders.Width())
}

// Update the system
func (s *SysInitLetters) Update(world *ecs.World) {
	if rand.Float64() > s.SpawnProb {
		return
	}

	grid := s.grid.Get()
	tick := s.time.Get().Tick
	letters := s.letters.Get().Letters

	e := s.builder.New()
	pos, let, mov := s.builder.Get(e)

	var py int
	for {
		py = rand.Intn(grid.Faders.Width())
		rel := s.releases[py]
		if rel == 0 || tick-rel > int64(s.MinGap) {
			break
		}
	}
	pos.X = py
	pos.Y = 0
	let.Letter = letters[rand.Intn(len(letters))]
	let.Size = rand.Intn(len(fontSizes))
	mov.LastMove = tick
	mov.Interval = uint16(rand.Intn(s.MaxMoveInterval-s.MinMoveInterval) + s.MinMoveInterval)
	mov.PathLength = rand.Intn(grid.Faders.Height())

	s.releases[pos.X] = tick
}

// Finalize the system
func (s *SysInitLetters) Finalize(world *ecs.World) {}
