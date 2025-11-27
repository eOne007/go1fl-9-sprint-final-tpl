package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateRandomElementsEmptyOrNonPositive(t *testing.T) {
	assert.Empty(t, generateRandomElements(-1000))
	assert.Empty(t, generateRandomElements(-1))
	assert.Empty(t, generateRandomElements(0))
}

func TestGenerateRandomElementsCorrectLength(t *testing.T) {
	assert.Len(t, generateRandomElements(1000), 1000)
}

func TestGenerateRandomElementsCorrectValues(t *testing.T) {
	data := generateRandomElements(50)
	for _, v := range data {
		assert.GreaterOrEqual(t, v, 0)
		assert.Less(t, v, 1000)
	}
}

func TestMaximumEmpty(t *testing.T) {
	assert.Equal(t, 0, maximum([]int{}))
}
func TestMaximumSingle(t *testing.T) {
	assert.Equal(t, 73, maximum([]int{73}))
}

func TestMaximumNegative(t *testing.T) {
	assert.Equal(t, -14, maximum([]int{-997, -14, -765}))
}

func TestMaximumAllCases(t *testing.T) {
	assert.Equal(t, 120, maximum([]int{-365, 0, 120}))
}

func TestMaxChunksEmpty(t *testing.T) {
	assert.Equal(t, 0, maxChunks([]int{}))
}

func TestMaxChunksSmallSlice(t *testing.T) {
	assert.Equal(t, 23, maxChunks([]int{-15, 0, 23}))
}

func TestMaxChunksMatchingResults(t *testing.T) {
	data := []int{-12, 0, 23, 456, -378, 120}
	assert.Equal(t, maximum(data), maxChunks(data))
}
