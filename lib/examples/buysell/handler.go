package buysell

import (
	"math/rand"

	"github.com/marcusljx/eevee3/lib/eevee3"
)

type Handler struct {
	Rand   *rand.Rand
	Prices []float64
}

func (h *Handler) randomAction() Action {
	return []Action{BUY, SELL, HOLD}[h.Rand.Int()%3]
}

func (h *Handler) Cross(solution1, solution2 eevee3.Solution[TUnderlying]) (eevee3.Solution[TUnderlying], eevee3.Solution[TUnderlying]) {
	var (
		items1 = solution1.Value()
		items2 = solution2.Value()
	)
	eevee3.CrossoverSliceByRandomPoint(items1, items2)
	return h.NewSolutionFrom(items1), h.NewSolutionFrom(items2)
}

func (h *Handler) Mutate(solution eevee3.Solution[TUnderlying]) eevee3.Solution[TUnderlying] {
	rosterCopy := solution.Value()
	eevee3.MutateRandomIndex(rosterCopy, func(_b Action) Action { return h.randomAction() })
	return h.NewSolutionFrom(rosterCopy)
}

func (h *Handler) NewSolution() eevee3.Solution[TUnderlying] {
	actions := make(TUnderlying, len(h.Prices))
	for i := range actions {
		actions[i] = h.randomAction()
	}
	return h.NewSolutionFrom(actions)
}

func (h *Handler) NewSolutionFrom(innerValue TUnderlying) eevee3.Solution[TUnderlying] {
	return &Solution{
		ConsistentSizeSliceSolution: *eevee3.NewConsistentSizeSliceSolution[Action](innerValue),
		handler:                     h,
	}
}
