package eevee3

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

var (
	rng = rand.New(rand.NewSource(time.Now().Unix()))
)

type SubgroupSelectionStrategy[T any] func(corpus []Solution[T], k int) []Solution[T]

func SelectRandomSubgroup[T any](corpus []Solution[T], k int) (result []Solution[T]) {
	if k > len(corpus) {
		panic("k cannot be larger than len(corpus)")
	}

	if k == len(corpus) {
		return corpus
	}

	indices := orderedSlice(len(corpus))
	rng.Shuffle(len(indices), func(i, j int) {
		indices[i], indices[j] = indices[j], indices[i]
	})

	for _, idx := range indices[:k] {
		result = append(result, corpus[idx])
	}
	return
}

func SelectBestSubgroup[T any](corpus []Solution[T], k int) (result []Solution[T]) {
	if k > len(corpus) {
		panic(fmt.Sprintf("k cannot be larger than len(corpus) [k=%d; len(corpus)=%d]\n", k, len(corpus)))
	}

	if k == len(corpus) {
		return corpus
	}

	indices := orderedSlice(len(corpus))
	sort.Slice(indices, func(i, j int) bool {
		return corpus[indices[i]].Score() > corpus[indices[j]].Score()
	})

	for _, idx := range indices[:k] {
		result = append(result, corpus[idx])
	}
	return
}

type PairwiseSelectionStrategy[T any] func(corpus []Solution[T], k int) [][2]Solution[T]

// SelectRandomPairs returns a k-size slice of randomly picked
// pairs of solutions
func SelectRandomPairs[T any](corpus []Solution[T], k int) [][2]Solution[T] {
	tmp := SelectRandomSubgroup(corpus, 2*k)
	result := make([][2]Solution[T], k)
	for i := 0; i < len(tmp); i += 2 {
		result[i/2] = [2]Solution[T]{
			tmp[i],
			tmp[i+1],
		}
	}
	return result
}

// SelectBestPairs returns a k-size slice of pairs of solutions
// whose scores are elitist, ie. Best and 2nd-best are paired,
// 3rd-best and 4th-best are paired.
func SelectBestPairs[T any](corpus []Solution[T], k int) [][2]Solution[T] {
	tmp := SelectBestSubgroup(corpus, 2*k)
	result := make([][2]Solution[T], k)
	for i := 0; i < len(tmp); i += 2 {
		result[i/2] = [2]Solution[T]{
			tmp[i],
			tmp[i+1],
		}
	}
	return result
}
