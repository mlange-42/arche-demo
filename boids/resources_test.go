package boids_test

import (
	"testing"

	"github.com/mlange-42/arche-demo/boids"
	"github.com/stretchr/testify/assert"
)

func TestGrid(t *testing.T) {
	grid := boids.NewGrid(128, 64, 16, 8)

	assert.Equal(t, 8, grid.Cols)
	assert.Equal(t, 4, grid.Rows)
}
