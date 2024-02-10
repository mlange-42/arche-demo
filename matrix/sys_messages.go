package matrix

import (
	"math/rand"

	"github.com/mlange-42/arche-demo/common"
	"github.com/mlange-42/arche-model/resource"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

var messages = []string{
	"Arche", "Arche", "Arche",
	"World", "World", "Entity", "Entity", "Query", "Query",
	"Mask", "Filter",
	"Next", "Get", "Alive", "Add", "Remove", "Exchange", "Relation",
}

// SysMessages is a system that places suplimal messages.
type SysMessages struct {
	Count    int
	Duration int

	time     generic.Resource[resource.Tick]
	messages generic.Resource[Messages]
	grid     generic.Resource[LetterGrid]

	filter    generic.Filter2[Position, LetterForcer]
	forcerMap generic.Map2[Position, LetterForcer]
	letterMap generic.Map1[ForcedLetter]

	toRemove []ecs.Entity
}

// Initialize the system
func (s *SysMessages) Initialize(world *ecs.World) {
	s.time = generic.NewResource[resource.Tick](world)
	s.messages = generic.NewResource[Messages](world)
	s.grid = generic.NewResource[LetterGrid](world)

	s.filter = *generic.NewFilter2[Position, LetterForcer]()
	s.forcerMap = generic.NewMap2[Position, LetterForcer](world)
	s.letterMap = generic.NewMap1[ForcedLetter](world)
}

// Update the system
func (s *SysMessages) Update(world *ecs.World) {
	tick := s.time.Get().Tick
	grid := s.grid.Get()
	messages := s.messages.Get()

	query := s.filter.Query(world)
	count := query.Count()
	for query.Next() {
		pos, forcer := query.Get()
		msg := messages.messages[forcer.Message]
		ln := len(msg)

		if forcer.TickDone >= 0 {
			if tick >= forcer.TickDone {
				for j := 0; j < ln; j++ {
					x := j + pos.X
					let := s.letterMap.Get(grid.Faders.Get(x, pos.Y))
					let.Active = false
				}
				s.toRemove = append(s.toRemove, query.Entity())
			}
			continue
		}

		done := true
		for j := 0; j < ln; j++ {
			x := j + pos.X
			let := s.letterMap.Get(grid.Faders.Get(x, pos.Y))
			if !let.Traversed {
				done = false
				break
			}
		}
		if done {
			forcer.TickDone = tick + int64(s.Duration)
		}
	}

	count -= len(s.toRemove)
	for _, e := range s.toRemove {
		world.RemoveEntity(e)
	}
	s.toRemove = s.toRemove[:0]

	if count < s.Count {
		s.createMessage(s.Count-count, &grid.Faders, messages)
	}
}

// Finalize the system
func (s *SysMessages) Finalize(world *ecs.World) {}

func (s *SysMessages) createMessage(count int, grid *common.Grid[ecs.Entity], messages *Messages) {
	query := s.forcerMap.NewBatchQ(count)
	for query.Next() {
		pos, forcer := query.Get()

		forcer.TickDone = -1
		forcer.Message = rand.Intn(len(messages.messages))
		msg := messages.messages[forcer.Message]
		ln := len(msg)

		sx, sy := grid.Width(), grid.Height()
		for i := 0; i < 10; i++ {
			pos.X, pos.Y = rand.Intn(sx-ln), rand.Intn(sy)

			ok := true
			for j := 0; j < ln; j++ {
				x := j + pos.X
				let := s.letterMap.Get(grid.Get(x, pos.Y))
				if let.Active {
					ok = false
					break
				}
			}
			if !ok {
				continue
			}

			for j := 0; j < ln; j++ {
				x := j + pos.X
				let := s.letterMap.Get(grid.Get(x, pos.Y))
				let.Letter = msg[j]
				let.Active = true
				let.Traversed = false
			}
			break
		}
	}
}
