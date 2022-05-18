package eevee3

type SolutionPredicate[T any] func(solution Solution[T]) (PredicateReason, bool)
type PopulationPredicate[T any] func([]Solution[T]) (PredicateReason, bool)

type PredicateReason string

const (
	_predicateFalse                            PredicateReason = ""
	PredicateReasonPopulationWasEmpty                          = "population is empty"
	PredicateReasonPopulationAllScoresSame                     = "all scores were equal"
	PredicateReasonPopulationAllSolutionsEqual                 = "all solutions were equal"
)

// TrueWhenAllScoresSame returns a [PopulationPredicate] that
// returns true when all solutions of the input population corpus
// have the same score
func TrueWhenAllScoresSame[T any]() PopulationPredicate[T] {
	return func(pop []Solution[T]) (PredicateReason, bool) {
		if len(pop) <= 0 {
			return PredicateReasonPopulationWasEmpty, true
		}
		x := pop[0]
		for _, y := range pop[1:] {
			if y.Score() != x.Score() {
				return _predicateFalse, false
			}
		}
		return PredicateReasonPopulationAllScoresSame, true
	}
}

// TrueWhenAllSolutionsEqual returns a [PopulationPredicate] that
// returns true when all solutions of the input population corpus
// are identical
func TrueWhenAllSolutionsEqual[T any]() PopulationPredicate[T] {
	return func(pop []Solution[T]) (PredicateReason, bool) {
		if len(pop) <= 0 {
			return PredicateReasonPopulationWasEmpty, true
		}
		x := pop[0]
		for _, y := range pop[1:] {
			if !x.Equals(y) {
				return _predicateFalse, false
			}
		}
		return PredicateReasonPopulationAllSolutionsEqual, true
	}
}
