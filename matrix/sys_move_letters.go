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
	MessageProb float64

	time      generic.Resource[resource.Tick]
	grid      generic.Resource[LetterGrid]
	messages  generic.Resource[Messages]
	canvas    generic.Resource[common.EbitenImage]
	letters   generic.Resource[Letters]
	filter    generic.Filter4[Position, Letter, Mover, Message]
	faderMap  generic.Map2[Letter, Fader]
	forcedMap generic.Map2[Fader, ForcedLetter]

	toRemove []ecs.Entity
}

// Initialize the system
func (s *SysMoveLetters) Initialize(world *ecs.World) {
	s.time = generic.NewResource[resource.Tick](world)
	s.grid = generic.NewResource[LetterGrid](world)
	s.messages = generic.NewResource[Messages](world)
	s.canvas = generic.NewResource[common.EbitenImage](world)
	s.letters = generic.NewResource[Letters](world)
	s.filter = *generic.NewFilter4[Position, Letter, Mover, Message]()
	s.faderMap = generic.NewMap2[Letter, Fader](world)
	s.forcedMap = generic.NewMap2[Fader, ForcedLetter](world)
}

// Update the system
func (s *SysMoveLetters) Update(world *ecs.World) {
	grid := s.grid.Get()
	tick := s.time.Get().Tick
	letters := s.letters.Get().Letters
	messages := s.messages.Get().messages

	query := s.filter.Query(world)

	for query.Next() {
		pos, let, mov, msg := query.Get()

		if tick-mov.LastMove < int64(mov.Interval) {
			continue
		}
		// Activate the fader on the cell we are about to leave
		eFader := grid.Faders.Get(pos.X, pos.Y)
		fLet, fad := s.faderMap.Get(eFader)
		fLet.Letter = let.Letter
		fad.Intensity = 1.0
		fad.NextChange = 0

		// Remove this mover
		if pos.Y+1 >= grid.Faders.Height() || pos.Y+1 >= mov.PathLength {
			s.toRemove = append(s.toRemove, query.Entity())
			continue
		}
		// Move down and select new random letter
		pos.Y += 1
		if msg.Message < 0 {
			if rand.Float64() < s.MessageProb {
				msg.Message = rand.Intn(len(messages))
				msg.Index = 0
			}
		}
		if msg.Message < 0 {
			let.Letter = letters[rand.Intn(len(letters))]
		} else {
			let.Letter = messages[msg.Message][msg.Index]
			msg.Index++
			if msg.Index >= len(messages[msg.Message]) {
				msg.Message = -1
			}
		}

		mov.LastMove = tick

		// De-activate the fader below
		eFaderNew := grid.Faders.Get(pos.X, pos.Y)
		fad, forced := s.forcedMap.Get(eFaderNew)
		fad.Intensity = 0

		if forced.Active {
			let.Letter = forced.Letter
			forced.Traversed = true
		}
	}

	for _, e := range s.toRemove {
		world.RemoveEntity(e)
	}

	s.toRemove = s.toRemove[:0]
}

// Finalize the system
func (s *SysMoveLetters) Finalize(world *ecs.World) {}
