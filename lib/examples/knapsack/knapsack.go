package knapsack

import (
	"github.com/marcusljx/eevee3/lib/eevee3"
	"math/rand"
	"strings"
	"time"
)

const WEIGHT_LIMIT = 1000

var rng = rand.New(rand.NewSource(time.Now().Unix()))

// Item defines a single knapsack item
type Item struct {
	Name          string
	WeightInGrams int
	Value         float64
}

// TUnderlying is the underlying type generic of
// a single solution in the experiment corpus
type TUnderlying = []Item

// Solution is just a type alias to make the Handler's code cleaner,
// you can have the type alias inline if you want
type Solution struct {
	items []Item
}

func (s *Solution) String() string {
	return strings.Join(sliceMap(s.items, func(item Item) string {
		return item.Name
	}), ", ")
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

	result := valueTotal / float64(weightTotal)
	if weightTotal > WEIGHT_LIMIT {
		return -result
	}
	return result
}

// Handler is the type that handles operations for the knapsack experiment.
// It also stores custom const data that is specific to the knapsack experiment.
type Handler struct {
	Rand *rand.Rand

	// itemUniverse represents all possible items
	// that can be put into the knapsack
	Items []Item
}

func (h *Handler) Cross(solution1, solution2 eevee3.Solution[TUnderlying]) [2]eevee3.Solution[TUnderlying] {
	var (
		items1   = solution1.OriginalValue()
		items2   = solution2.OriginalValue()
		splitIdx = rng.Int() % min[int](len(items1), len(items2))
	)

	for i := 0; i < splitIdx; i++ {
		items1[i], items2[i] = items2[i], items1[i]
	}

	return [2]eevee3.Solution[TUnderlying]{
		&Solution{items: items1},
		&Solution{items: items2},
	}
}

func (h *Handler) Mutate(solution eevee3.Solution[TUnderlying]) eevee3.Solution[TUnderlying] {
	var (
		items     = solution.OriginalValue()
		targetIdx = rng.Int() % len(items)
		newIdx    = rng.Int() % len(h.Items)
	)

	items[targetIdx] = h.Items[newIdx]
	return &Solution{items: items}
}

func (h *Handler) NewSolution() eevee3.Solution[TUnderlying] {
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
