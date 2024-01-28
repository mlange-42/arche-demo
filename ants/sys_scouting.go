package ants

import (
	"math"
	"math/rand"

	"github.com/mlange-42/arche-model/resource"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

type returnEntry struct {
	Entity ecs.Entity
	Load   float64
}

// SysScouting is a system that performs scout decisions.
type SysScouting struct {
	MaxCollect    float64
	TraceDecay    float64
	MaxSearchTime int

	time        generic.Resource[resource.Tick]
	filter      generic.Filter2[AntEdge, ActScout]
	antEdgeMap  generic.Map1[AntEdge]
	returnMap   generic.Map1[ActReturn]
	edgeMap     generic.Map2[Edge, Trace]
	edgeGeomMap generic.Map1[EdgeGeometry]
	nodeMap     generic.Map1[Node]
	resourceMap generic.Map[NodeResource]

	exchangeReturn generic.Exchange

	toReturn []returnEntry
}

// Initialize the system
func (s *SysScouting) Initialize(world *ecs.World) {
	s.time = generic.NewResource[resource.Tick](world)
	s.filter = *generic.NewFilter2[AntEdge, ActScout]()

	s.antEdgeMap = generic.NewMap1[AntEdge](world)
	s.returnMap = generic.NewMap1[ActReturn](world)
	s.edgeMap = generic.NewMap2[Edge, Trace](world)
	s.edgeGeomMap = generic.NewMap1[EdgeGeometry](world)
	s.nodeMap = generic.NewMap1[Node](world)
	s.resourceMap = generic.NewMap[NodeResource](world)

	s.exchangeReturn = *generic.NewExchange(world).
		Adds(generic.T[ActReturn]()).
		Removes(generic.T[ActScout]())

	s.toReturn = make([]returnEntry, 0, 16)
}

// Update the system
func (s *SysScouting) Update(world *ecs.World) {
	tick := s.time.Get().Tick

	query := s.filter.Query(world)
	for query.Next() {
		antEdge, scout := query.Get()
		if antEdge.Pos < antEdge.Length {
			continue
		}
		oldEdge, oldTrace := s.edgeMap.Get(antEdge.Edge)
		oldTrace.FromNest += math.Pow(s.TraceDecay, float64(tick-scout.Start))

		if s.resourceMap.Has(oldEdge.To) {
			res := s.resourceMap.Get(oldEdge.To)

			entry := returnEntry{Entity: query.Entity(), Load: res.Resource}
			res.Resource -= s.MaxCollect
			if res.Resource < 0 {
				res.Resource = 0
			}
			s.toReturn = append(s.toReturn, entry)
			continue
		}

		if tick > scout.Start+int64(s.MaxSearchTime) {
			s.toReturn = append(s.toReturn, returnEntry{Entity: query.Entity(), Load: 0})
			continue
		}

		node := s.nodeMap.Get(oldEdge.To)
		for {
			idx := rand.Intn(node.NumEdges)
			if node.InEdges[idx] == antEdge.Edge {
				// Don't allow scouts to go where they came from.
				continue
			}
			edge := node.OutEdges[idx]
			geom := s.edgeGeomMap.Get(edge)
			antEdge.Update(edge, geom)
			break
		}
	}

	for _, e := range s.toReturn {
		s.exchangeReturn.Exchange(e.Entity)
		ret := s.returnMap.Get(e.Entity)
		ret.Start = tick
		ret.Load = e.Load
	}

	s.toReturn = s.toReturn[:0]
}

// Finalize the system
func (s *SysScouting) Finalize(world *ecs.World) {}
