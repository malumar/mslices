package mslices

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnique(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3, 4, 5}, Unique([]int{1, 2, 2, 3, 4, 5}))
	assert.Equal(t, []int{1, 2, 3, 4}, Unique([]int{1, 2, 2, 3, 4, 2}))
	assert.Equal(t, []int{1}, Unique([]int{1}))
	assert.NotEqual(t, []int{1, 2}, Unique([]int{1}))
	assert.NotEqual(t, nil, Unique([]int{1}))
	assert.Equal(t, nil, nil)
	assert.NotEqual(t, []int{}, nil)
	assert.NotEqual(t, []int{1}, nil)
	assert.NotEqual(t, []int{1}, Unique([]int{}))
	assert.Equal(t, []int{}, Unique([]int{}))
}

func TestFilter(t *testing.T) {
	var empty []int
	var empty2 []int
	assert.Equal(t, empty2, Filter(empty, func(val int) bool {
		return true
	}))
	assert.Equal(t, []int{2, 4, 6}, Filter([]int{1, 2, 3, 4, 5, 6}, func(val int) bool {
		return val%2 == 0
	}))
}
