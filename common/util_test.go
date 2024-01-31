package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandBetweenUInt8(t *testing.T) {
	counts := make([]int, 10)
	for i := 0; i < 1000; i++ {
		counts[RandBetweenUIn8(5, 3)]++
	}

	assert.Equal(t, 0, counts[2])
	assert.NotEqual(t, 0, counts[3])
	assert.NotEqual(t, 0, counts[4])
	assert.NotEqual(t, 0, counts[5])
	assert.Equal(t, 0, counts[6])
}
