package eevee3

import (
	"math/rand"
	"time"
)

var (
	rng = rand.New(rand.NewSource(time.Now().Unix()))
)

// orderedSlice returns a slice of
// incrementing values (starting from 0)
func orderedSlice(n int) []int {
	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = i
	}
	return result
}

func min[T int | float64](val T, values ...T) T {
	for _, y := range values {
		if y < val {
			val = y
		}
	}
	return val
}
