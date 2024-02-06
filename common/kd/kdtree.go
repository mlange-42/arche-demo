package kd

import (
	"container/heap"

	"github.com/mlange-42/arche-demo/common"
	"github.com/mlange-42/arche/ecs"
	"gonum.org/v1/gonum/spatial/kdtree"
)

type EntityLocation struct {
	common.Vec2f
	Heading float64
	Entity  ecs.Entity
}

// Compare returns the signed distance of p from the plane passing through c and
// perpendicular to the dimension d. The concrete type of c must be EntityLocation.
func (p EntityLocation) Compare(c kdtree.Comparable, d kdtree.Dim) float64 {
	q := c.(EntityLocation)
	return p.Get(int(d)) - q.Get(int(d))
}

// Dims returns the number of dimensions described by the receiver.
func (p EntityLocation) Dims() int { return 2 }

// Distance returns the squared Euclidean distance between c and the receiver. The
// concrete type of c must be EntityLocation.
func (p EntityLocation) Distance(c kdtree.Comparable) float64 {
	q := c.(EntityLocation)
	dx := q.X - p.X
	dy := q.Y - p.Y
	return dx*dx + dy*dy
}

type EntityLocations []EntityLocation

func (p EntityLocations) Index(i int) kdtree.Comparable { return p[i] }
func (p EntityLocations) Len() int                      { return len(p) }
func (p EntityLocations) Pivot(d kdtree.Dim) int {
	return plane{EntityLocations: p, Dim: d}.Pivot()
}
func (p EntityLocations) Slice(start, end int) kdtree.Interface { return p[start:end] }

// plane is a wrapping type that allows a Points type be pivoted on a dimension.
// The Pivot method of Plane uses MedianOfRandoms sampling at most 100 elements
// to find a pivot element.
type plane struct {
	kdtree.Dim
	EntityLocations
}

// randoms is the maximum number of random values to sample for calculation of
// median of random elements.
const randoms = 100

// Less comparison
func (p plane) Less(i, j int) bool {
	return p.EntityLocations[i].Get(int(p.Dim)) < p.EntityLocations[j].Get(int(p.Dim))
}

// Pivot TreePlane
func (p plane) Pivot() int { return kdtree.Partition(p, kdtree.MedianOfRandoms(p, randoms)) }

// Slice TreePlane
func (p plane) Slice(start, end int) kdtree.SortSlicer {
	p.EntityLocations = p.EntityLocations[start:end]
	return p
}

// Swap TreePlane
func (p plane) Swap(i, j int) {
	p.EntityLocations[i], p.EntityLocations[j] = p.EntityLocations[j], p.EntityLocations[i]
}

// NDistKeeper keeps man number and distance
type NDistKeeper struct {
	kdtree.Heap
}

// NewNDistKeeper returns an NDistKeeper with the maximum value of the heap set to d.
func NewNDistKeeper(n int, d float64) *NDistKeeper {
	k := NDistKeeper{make(kdtree.Heap, 1, n)}
	k.Heap[0].Dist = d * d
	return &k
}

// Keep adds c to the heap if its distance is less than or equal to the max value of the heap.
func (k *NDistKeeper) Keep(c kdtree.ComparableDist) {
	if c.Dist <= k.Heap[0].Dist { // Favour later finds to displace sentinel.
		if len(k.Heap) == cap(k.Heap) {
			heap.Pop(k)
		}
		heap.Push(k, c)
	}
}
