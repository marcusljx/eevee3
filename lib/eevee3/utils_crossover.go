package eevee3

import "fmt"

// CrossoverSliceByMidpoint performs an in-place crossover
// on an exact mid-point split of both slices.
// It is the caller's responsibility to ensure that
// the input slices are safe to be mutated.
//
// This function panics if len(s1) != len(s2)
func CrossoverSliceByMidpoint[T any](s1, s2 []T) {
	if len(s1) != len(s2) {
		panic(fmt.Sprintf("expected len(s1)==len(s2), but got len(s1)=%d, len(s2)=%d", len(s1), len(s2)))
	}

	for i := 0; i < (len(s1) / 2); i++ {
		s1[i], s2[i] = s2[i], s1[i]
	}
}

// CrossoverSliceByRandomPoint performs an in-place crossover
// on a random point-of-split of both slices.
// The point-of-split is randomly decided,
// but is guaranteed to be within the bounds of both slices.
// It is the caller's responsibility to ensure that
// the input slices are safe to be mutated.
//
// This function panics if len(s1) != len(s2)
func CrossoverSliceByRandomPoint[T any](s1, s2 []T) {
	maxSplitIdx := min[int](len(s1), len(s2))
	if maxSplitIdx == 0 {
		return
	}

	splitIdx := rng.Int() % maxSplitIdx
	for i := 0; i < splitIdx; i++ {
		s1[i], s2[i] = s2[i], s1[i]
	}
}
