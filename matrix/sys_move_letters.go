package matrix

import (
	"math/rand"

	"github.com/mlange-42/arche-demo/common"
	"github.com/mlange-42/arche-model/resource"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// SysMoveLetters is a system to move letter and activate/de-activate faders.
type SysMoveLetters struct {
	time     generic.Resource[resource.Tick]
	grid     generic.Resource[LetterGrid]
	canvas   generic.Resource[common.EbitenImage]
	letters  generic.Resource[Letters]
	filter   generic.Filter3[Position, Letter, Mover]
	faderMap generic.Map2[Letter, Fader]

	toRemove []ecs.Entity
}

// Initialize the system
func (s *SysMoveLetters) Initialize(world *ecs.World) {
	s.time = generic.NewResource[resource.Tick](world)
	s.grid = generic.NewResource[LetterGrid](world)
	s.canvas = generic.NewResource[common.EbitenImage](world)
	s.letters = generic.NewResource[Letters](world)
	s.filter = *generic.NewFilter3[Position, Letter, Mover]()
	s.faderMap = generic.NewMap2[Letter, Fader](world)
}

// Update the system
func (s *SysMoveLetters) Update(world *ecs.World) {
	grid := s.grid.Get()
	tick := s.time.Get().Tick
	letters := s.letters.Get().Letters

	query := s.filter.Query(world)
	for query.Next() {
		pos, let, mov := query.Get()

		if tick-mov.LastMove < int64(mov.Interval) {
			continue
		}
		// Activate the fader on the cell we are about to leave
		eFader := grid.Faders.Get(pos.X, pos.Y)
		fLet, fad := s.faderMap.Get(eFader)
		fLet.Letter = let.Letter
		fad.Intensity = 1.0

		// Remove this mover
		if pos.Y+1 >= grid.Faders.Height() || pos.Y+1 >= mov.PathLength {
			s.toRemove = append(s.toRemove, query.Entity())
			continue
		}
		// Move down and select new random letter
		pos.Y += 1
		let.Letter = letters[rand.Intn(len(letters))]
		mov.LastMove = tick

		// De-activate the fader below
		eFaderNew := grid.Faders.Get(pos.X, pos.Y)
		_, fad = s.faderMap.Get(eFaderNew)
		fad.Intensity = 0
	}

	for _, e := range s.toRemove {
		world.RemoveEntity(e)
	}

	s.toRemove = s.toRemove[:0]
}

// Finalize the system
func (s *SysMoveLetters) Finalize(world *ecs.World) {}
