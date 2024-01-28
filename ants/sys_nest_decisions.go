package ants

import (
	"math/rand"

	"github.com/mlange-42/arche-model/resource"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// SysNestDecisions is a system that performs ant decisions in the nest.
//
// Particularly, it decides whether bees go for scouting or foraging.
type SysNestDecisions struct {
	ReleaseInterval  int64
	ReleaseCount     int
	ScoutProbability float64

	time       generic.Resource[resource.Tick]
	nest       generic.Resource[Nest]
	idleFilter generic.Filter0

	exchangeScout  generic.Exchange
	scoutMap       generic.Map1[ActScout]
	exchangeForage generic.Exchange

	antEdgeMap generic.Map1[AntEdge]
	nodeMap    generic.Map1[Node]
	edgeMap    generic.Map1[EdgeGeometry]
	traceMap   generic.Map1[Trace]

	toLeave []ecs.Entity
}

type waggleInfo struct {
	Target  Position
	Benefit float64
}

// Initialize the system
func (s *SysNestDecisions) Initialize(world *ecs.World) {
	s.time = generic.NewResource[resource.Tick](world)
	s.nest = generic.NewResource[Nest](world)
	s.idleFilter = *generic.NewFilter0().With(generic.T[ActInNest]())

	s.exchangeScout = *generic.NewExchange(world).
		Adds(generic.T2[ActScout, AntEdge]()...).
		Removes(generic.T[ActInNest]())
	s.scoutMap = generic.NewMap1[ActScout](world)

	s.exchangeForage = *generic.NewExchange(world).
		Adds(generic.T2[ActForage, AntEdge]()...).
		Removes(generic.T[ActInNest]())

	s.antEdgeMap = generic.NewMap1[AntEdge](world)
	s.nodeMap = generic.NewMap1[Node](world)
	s.edgeMap = generic.NewMap1[EdgeGeometry](world)
	s.traceMap = generic.NewMap1[Trace](world)

	s.toLeave = make([]ecs.Entity, 0, 64)
}

// Update the system
func (s *SysNestDecisions) Update(world *ecs.World) {
	tick := s.time.Get().Tick

	idleQuery := s.idleFilter.Query(world)
	totalCnt := idleQuery.Count()
	cnt := 0
	for idleQuery.Next() {
		if cnt >= s.ReleaseCount {
			break
		}
		s.toLeave = append(s.toLeave, idleQuery.Entity())
		cnt++
	}
	if cnt < totalCnt {
		idleQuery.Close()
	}

	nest := s.nest.Get()
	node := s.nodeMap.Get(nest.Node)

	for _, e := range s.toLeave {
		if rand.Float64() < s.ScoutProbability {
			s.exchangeScout.Exchange(e)
			scout := s.scoutMap.Get(e)
			scout.Start = tick

			edge := node.OutEdges[rand.Intn(node.NumEdges)]
			geom := s.edgeMap.Get(edge)

			ant := s.antEdgeMap.Get(e)
			ant.Update(edge, geom)
			continue
		}

		s.exchangeForage.Exchange(e)
		edge := node.OutEdges[rand.Intn(node.NumEdges)]
		geom := s.edgeMap.Get(edge)

		ant := s.antEdgeMap.Get(e)
		ant.Edge = edge
		ant.From = geom.From
		ant.Dir = geom.Dir
		ant.Length = geom.Length
		ant.Pos = 0
	}

	s.toLeave = s.toLeave[:0]
}

func (s *SysNestDecisions) selectEdge(world *ecs.World, node *Node) ecs.Entity {
	maxTrace := -1.0
	var maxEdge ecs.Entity
	count := 0
	for i := 0; i < node.NumEdges; i++ {
		edge := node.InEdges[i]
		trace := s.traceMap.Get(edge)
		if trace.FromResource > maxTrace {
			maxTrace = trace.FromResource
			maxEdge = edge
			count = 1
			continue
		}
		if trace.FromResource == maxTrace {
			count++
			if rand.Float64() < 1.0/float64(count) {
				maxEdge = edge
			}
		}
	}
	return maxEdge
}

// Finalize the system
func (s *SysNestDecisions) Finalize(world *ecs.World) {}
