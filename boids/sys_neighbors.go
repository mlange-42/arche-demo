package boids

import (
	"github.com/mlange-42/arche-demo/common/kd"
	"github.com/mlange-42/arche-model/resource"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
	"gonum.org/v1/gonum/spatial/kdtree"
)

// SysNeighbors system.
type SysNeighbors struct {
	Neighbors      int
	Radius         float64
	BuildStep      int
	posFilter      generic.Filter1[Position]
	posNeighFilter generic.Filter3[Position, Velocity, Neighbors]
	tickRes        generic.Resource[resource.Tick]

	points []kd.EntityLocation
}

// Initialize the system
func (s *SysNeighbors) Initialize(w *ecs.World) {
	if s.Neighbors > MaxNeighbors {
		panic("maximum number of neighbors exceeded. See constant MaxNeighbors")
	}
	s.posFilter = *generic.NewFilter1[Position]()
	s.posNeighFilter = *generic.NewFilter3[Position, Velocity, Neighbors]()

	s.posFilter.Register(w)
	s.posNeighFilter.Register(w)

	s.tickRes = generic.NewResource[resource.Tick](w)
}

// Update the system
func (s *SysNeighbors) Update(w *ecs.World) {
	if s.BuildStep > 0 && s.tickRes.Get().Tick%int64(s.BuildStep) != 0 {
		return
	}
	tree := s.createTree(w)

	query := s.posNeighFilter.Query(w)

	for query.Next() {
		entity := query.Entity()
		pos, vel, neigh := query.Get()

		p := kd.EntityLocation{Vec2f: pos.Vec2f, Velocity: vel.Vec2f, Entity: entity}
		keep := kd.NewNDistKeeper(s.Neighbors+1, s.Radius)
		tree.NearestSet(keep, p)

		neigh.Count = 0

		if keep.Heap.Len() > 1 {
			for _, c := range keep.Heap {
				n := c.Comparable.(kd.EntityLocation)
				if n.Entity == entity {
					continue
				}
				neigh.Entities[neigh.Count] = n
				neigh.Count++
			}
		}
	}
}

// Finalize the system
func (s *SysNeighbors) Finalize(w *ecs.World) {}

func (s *SysNeighbors) createTree(w *ecs.World) *kdtree.Tree {
	query := s.posFilter.Query(w)
	for query.Next() {
		e := query.Entity()
		pos := query.Get()
		s.points = append(s.points, kd.EntityLocation{Vec2f: pos.Vec2f, Entity: e})
	}
	treePoints := kd.EntityLocations(s.points)
	tree := kdtree.New(treePoints, false)
	s.points = s.points[:0]
	return tree
}
