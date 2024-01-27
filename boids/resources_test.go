package boids_test

import (
	"testing"

	"github.com/mlange-42/arche-demo/boids"
	"github.com/mlange-42/arche/ecs"
	"github.com/stretchr/testify/assert"
)

func TestGrid(t *testing.T) {
	grid := boids.NewGrid(128, 64, 16, 8)

	assert.Equal(t, 8, grid.Cols)
	assert.Equal(t, 4, grid.Rows)

	pos := boids.Position{X: 1, Y: 1}
	vel := boids.Velocity{X: 2, Y: 2}

	grid.Add(1, 2, ecs.Entity{}, &pos, &vel)

	assert.Equal(t, []boids.GridEntry{{ecs.Entity{}, pos.X, pos.Y, vel.X, vel.Y}}, grid.Get(1, 2))
	assert.Equal(t, []boids.GridEntry{}, grid.Get(2, 1))

	grid.Add(1, 2, ecs.Entity{}, &pos, &vel)
	grid.Add(1, 2, ecs.Entity{}, &pos, &vel)
	grid.Add(1, 2, ecs.Entity{}, &pos, &vel)
	grid.Add(1, 2, ecs.Entity{}, &pos, &vel)
	grid.Add(1, 2, ecs.Entity{}, &pos, &vel)
	grid.Add(1, 2, ecs.Entity{}, &pos, &vel)
	grid.Add(1, 2, ecs.Entity{}, &pos, &vel)

	assert.Panics(t, func() {
		grid.Add(1, 2, ecs.Entity{}, &pos, &vel)
	})

	assert.Equal(t, 8, len(grid.Get(1, 2)))
	assert.Equal(t, 0, len(grid.Get(1, 3)))

	grid.Clear()
	assert.Equal(t, 0, len(grid.Get(1, 2)))
}
