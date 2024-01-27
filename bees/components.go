package bees

import (
	"github.com/mlange-42/arche/ecs"
)

// Position component.
type Position struct {
	X float64
	Y float64
}

// Direction component.
type Direction struct {
	X float64
	Y float64
}

// HomeHive component. This is an [ecs.Relation] component,
// associating each bee with it's home hive.
type HomeHive struct {
	ecs.Relation
}

// Random256 component, contains an uint8 value for scheduling things in intervals, but randomized over entities.
type Random256 struct {
	Value uint8
}

// Hive component, as a label for bee hives.
type Hive struct{}

// FlowerPatch component.
type FlowerPatch struct {
	X         int
	Y         int
	Resources float64
}
