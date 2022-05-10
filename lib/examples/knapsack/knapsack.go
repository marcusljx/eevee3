package knapsack

import (
	"github.com/marcusljx/eevee3/lib/eevee3"
	"math/rand"
)

// Item defines a single knapsack item
type Item struct {
	WeightInGrams int
	Value         float64
}

// SolutionType is the underlying type generic of
// a single solution in the experiment corpus
type SolutionType = []Item

// Solution is just a type alias to make the Handler's code cleaner,
// you can have the type alias inline if you want
type Solution struct {
	items []Item
}

func (s *Solution) OriginalValue() []Item {
	return s.items
}

// Score for knapsack is the value-per-gram
func (s *Solution) Score() float64 {
	weightTotal := sliceFold[Item, int](0, s.items, func(current int, next Item) int {
		return current + next.WeightInGrams
	})

	valueTotal := sliceFold[Item, float64](0, s.items, func(current float64, next Item) float64 {
		return current + next.Value
	})

	return valueTotal / float64(weightTotal)
}

// Handler is the type that handles operations for the knapsack experiment.
// It also stores custom const data that is specific to the knapsack experiment.
type Handler struct {
	Rand *rand.Rand

	// itemUniverse represents all possible items
	// that can be put into the knapsack
	Items []Item
}

func (h *Handler) NewSolution() eevee3.Solution[[]Item] {
	var (
		count    = h.Rand.Int() % len(h.Items)
		selected = make([]Item, count)
	)
	h.Rand.Shuffle(len(h.Items), func(i, j int) {
		h.Items[i], h.Items[j] = h.Items[j], h.Items[i]
	})
	copy(selected, h.Items)

	return &Solution{
		items: selected,
	}
}

func (h *Handler) Cross(solution1, solution2 eevee3.Solution[[]Item]) [2]eevee3.Solution[[]Item] {
	//TODO implement me
	panic("implement me")
}

func (h *Handler) Mutate(solution eevee3.Solution[[]Item]) eevee3.Solution[[]Item] {
	//TODO implement me
	panic("implement me")
}
