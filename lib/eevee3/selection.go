package eevee3

import (
	"math"
	"sort"
)

type SubgroupSelectionStrategy[T any] func(corpus []Solution[T]) []Solution[T]

func SelectRandomSubgroup[T any]() SubgroupSelectionStrategy[T] {
	return func(corpus []Solution[T]) (result []Solution[T]) {
		n := len(corpus)

		indices := orderedSlice(len(corpus))
		rng.Shuffle(len(indices), func(i, j int) {
			indices[i], indices[j] = indices[j], indices[i]
		})

		for _, idx := range indices[:n] {
			result = append(result, corpus[idx])
		}
		return
	}
}

func SelectBestSubgroup[T any]() SubgroupSelectionStrategy[T] {
	return func(corpus []Solution[T]) (result []Solution[T]) {
		n := len(corpus)

		indices := orderedSlice(len(corpus))
		sort.Slice(indices, func(i, j int) bool {
			return corpus[indices[i]].Score() > corpus[indices[j]].Score()
		})

		for _, idx := range indices[:n] {
			result = append(result, corpus[idx])
		}
		return
	}
}

func SelectBestAndWorstSubgroup[T any](bestRatio float64) SubgroupSelectionStrategy[T] {
	return func(corpus []Solution[T]) (result []Solution[T]) {
		var (
			n      = len(corpus)
			kBest  = int(math.Floor(bestRatio * float64(n)))
			kWorst = n - kBest
		)

		indices := orderedSlice(len(corpus))
		sort.Slice(indices, func(i, j int) bool {
			return corpus[indices[i]].Score() > corpus[indices[j]].Score()
		})

		for _, idx := range indices[:kBest] {
			result = append(result, corpus[idx])
		}
		for _, idx := range indices[len(indices)-kWorst:] {
			result = append(result, corpus[idx])
		}

		return
	}
}

type PairwiseSelectionStrategy[T any] func(corpus []Solution[T]) [][2]Solution[T]

// SelectRandomPairs returns a k-size slice of randomly picked
// pairs of solutions
func SelectRandomPairs[T any]() PairwiseSelectionStrategy[T] {
	return func(corpus []Solution[T]) [][2]Solution[T] {
		var (
			n      = len(corpus)
			tmp    = SelectRandomSubgroup[T]()(corpus)
			result = make([][2]Solution[T], n/2)
		)
		for i := 0; i < len(tmp); i += 2 {
			result[i/2] = [2]Solution[T]{
				tmp[i],
				tmp[i+1],
			}
		}
		return result
	}
}

// SelectBestPairs returns a k-size slice of pairs of solutions
// whose scores are elitist, ie. Best and 2nd-best are paired,
// 3rd-best and 4th-best are paired.
func SelectBestPairs[T any]() PairwiseSelectionStrategy[T] {
	return func(corpus []Solution[T]) [][2]Solution[T] {
		var (
			n      = len(corpus)
			tmp    = SelectBestSubgroup[T]()(corpus)
			result = make([][2]Solution[T], n/2)
		)
		for i := 0; i < len(tmp); i += 2 {
			result[i/2] = [2]Solution[T]{
				tmp[i],
				tmp[i+1],
			}
		}
		return result
	}
}
