package knapsack

import (
	"github.com/marcusljx/eevee3/lib/eevee3"
	"math/rand"
	"time"
)

var rng = rand.New(rand.NewSource(time.Now().Unix()))

// Item defines a single knapsack item
type Item struct {
	Name          string
	WeightInGrams int
	Value         float64
}

// Handler is the type that handles operations for the knapsack experiment.
// It also stores custom const data that is specific to the knapsack experiment.
type Handler struct {
	Rand *rand.Rand

	// itemUniverse represents all possible items
	// that can be put into the knapsack
	Items []Item
}

func (h *Handler) Cross(solution1, solution2 eevee3.Solution[TUnderlying]) (eevee3.Solution[TUnderlying], eevee3.Solution[TUnderlying]) {
	var (
		items1 = solution1.Value()
		items2 = solution2.Value()
	)
	eevee3.CrossoverSliceByRandomPoint(items1, items2)
	return h.newSolution(items1), h.newSolution(items2)
}

func (h *Handler) Mutate(solution eevee3.Solution[TUnderlying]) eevee3.Solution[TUnderlying] {
	rosterCopy := solution.Value()
	eevee3.MutateRandomIndex(rosterCopy, func(b bool) bool { return !b })
	return h.newSolution(rosterCopy)
}

func (h *Handler) NewSolution() eevee3.Solution[TUnderlying] {
	roster := make([]bool, len(h.Items))
	for i := range roster {
		roster[i] = rng.Float64() < 0.5
	}
	return h.newSolution(roster)
}

func (h *Handler) newSolution(roster []bool) *Solution {
	return &Solution{
		handler:                     h,
		ConsistentSizeSliceSolution: *eevee3.NewConsistentSizeSliceSolution[bool](roster),
	}
}
