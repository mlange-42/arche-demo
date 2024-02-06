package matrix

import (
	"math/rand"

	"github.com/mlange-42/arche-demo/common"
	"github.com/mlange-42/arche-model/resource"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// SysMoveLetters is a system to initialize entities.
type SysMoveLetters struct {
	time    generic.Resource[resource.Tick]
	grid    generic.Resource[LetterGrid]
	canvas  generic.Resource[common.EbitenImage]
	letters generic.Resource[Letters]
	filter  generic.Filter3[Position, Letter, Mover]

	builder generic.Map3[Position, Letter, Fader]

	toRemove []ecs.Entity
	toSpawn  []spawnEntry
}

type spawnEntry struct {
	Pos    Position
	Letter rune
}

// Initialize the system
func (s *SysMoveLetters) Initialize(world *ecs.World) {
	s.time = generic.NewResource[resource.Tick](world)
	s.grid = generic.NewResource[LetterGrid](world)
	s.canvas = generic.NewResource[common.EbitenImage](world)
	s.letters = generic.NewResource[Letters](world)
	s.filter = *generic.NewFilter3[Position, Letter, Mover]()
	s.builder = generic.NewMap3[Position, Letter, Fader](world)
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
		s.toSpawn = append(s.toSpawn, spawnEntry{*pos, let.Letter})

		if pos.Y+1 >= grid.Faders.Height() || pos.Y+1 >= mov.PathLength {
			s.toRemove = append(s.toRemove, query.Entity())
			continue
		}
		pos.Y += 1
		let.Letter = letters[rand.Intn(len(letters))]
		mov.LastMove = tick

		fader := grid.Faders.Get(pos.X, pos.Y)
		if fader.IsZero() {
			continue
		}
		s.toRemove = append(s.toRemove, fader)
	}

	for _, e := range s.toRemove {
		if world.Alive(e) {
			world.RemoveEntity(e)
		}
	}
	s.toRemove = s.toRemove[:0]

	if len(s.toSpawn) == 0 {
		return
	}
	querySpawn := s.builder.NewBatchQ(len(s.toSpawn))
	cnt := 0
	for querySpawn.Next() {
		entry := s.toSpawn[cnt]
		pos, let, fad := querySpawn.Get()

		*pos = entry.Pos
		let.Letter = entry.Letter
		fad.Start = tick
		fad.Intensity = 1.0

		cnt++
	}
	s.toSpawn = s.toSpawn[:0]
}

// Finalize the system
func (s *SysMoveLetters) Finalize(world *ecs.World) {}
