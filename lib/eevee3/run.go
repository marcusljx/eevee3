package eevee3

import (
	"fmt"
	"log"
	"sort"
	"strings"
)

const (
	defaultTerminationReason = "reached specified max generations"
)

// RunN is similar to Run, but only returns the top (by score)
// n solutions in the final population corpus
func RunN[T any](handler Handler[T], controller *Controller[T], n int) *Result[T] {
	result := Run[T](handler, controller)
	sort.Slice(result.Population, func(i, j int) bool {
		return result.Population[i].Score() > result.Population[j].Score()
	})
	result.Population = result.Population[:n]

	return result
}

// RunSingle is a helper for RunN(handler, controller, 1)
func RunSingle[T any](handler Handler[T], controller *Controller[T]) *Result[T] {
	return RunN(handler, controller, 1)
}

func Run[T any](handler Handler[T], controller *Controller[T]) *Result[T] {
	var (
		pop        = createPopulation(handler, controller)
		generation = 1
	)

	logPopulation(0, pop)
	for ; generation <= controller.GenerationCycles; generation++ {
		iterate(pop, handler, controller)
		logPopulation(generation, pop)

		// evaluate termination conditions
		for _, cond := range controller.TerminationConditions {
			if reason, stop := cond(pop); stop {
				return &Result[T]{
					Population:             pop,
					TerminatedAtGeneration: generation,
					TerminationReason:      string(reason),
				}
			}
		}
	}
	return &Result[T]{
		Population:             pop,
		TerminatedAtGeneration: generation,
		TerminationReason:      defaultTerminationReason,
	}
}

func logPopulation[T any](generation int, pop []Solution[T]) {
	scores := make([]string, len(pop))
	for i, sol := range pop {
		scores[i] = fmt.Sprintf("%+0.2f", sol.Score())
	}
	log.Printf("[gen:%04d][%s]", generation, strings.Join(scores, ","))
}

func createPopulation[T any](handler Handler[T], data *Controller[T]) []Solution[T] {
	result := make([]Solution[T], data.PopulationSize)
	for i := 0; i < data.PopulationSize; i++ {
		result[i] = handler.NewSolution()
	}
	return result
}

func iterate[T any](pop []Solution[T], handler Handler[T], data *Controller[T]) {
	var (
		crossoverInChan         = make(chan [2]Solution[T])
		crossoverToMutationChan = make(chan Solution[T])
		out                     = make(chan Solution[T])
	)

	go doCrossover[T](crossoverInChan, crossoverToMutationChan, data.CrossoverProbability, handler.Cross)
	go doMutation[T](crossoverToMutationChan, out, data.MutationProbability, handler.Mutate)

	// feed in elements
	go func() {
		for _, pair := range data.CrossoverSelectionStrategy(pop) {
			crossoverInChan <- pair
		}
		close(crossoverInChan)
	}()

	// read out
	populationPool := pop
	for sol := range out {
		populationPool = append(populationPool, sol)
	}
	//logPopulation(-999, populationPool)

	copy(pop, data.NextGenerationSelectionStrategy(populationPool))
}

func doCrossover[T any](
	in chan [2]Solution[T],
	out chan Solution[T],
	probability float64,
	crossoverFunc func(s1, s2 Solution[T]) (Solution[T], Solution[T]),
) {
	for pair := range in {
		result1, result2 := pair[0], pair[1]
		if rng.Float64() < probability {
			result1, result2 = crossoverFunc(pair[0], pair[1])
		}
		out <- result1
		out <- result2
	}
	close(out)
}

func doMutation[T any](
	in chan Solution[T],
	out chan Solution[T],
	probability float64,
	mutateFunc func(s1 Solution[T]) Solution[T],
) {
	for sol := range in {
		result := sol
		if rng.Float64() < probability {
			result = mutateFunc(sol)
		}
		out <- result
	}
	close(out)
}
