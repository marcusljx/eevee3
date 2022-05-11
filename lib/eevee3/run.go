package eevee3

func Run[T any](handler Handler[T], data *ExperimentData[T]) []Solution[T] {
	pop := createPopulation(handler, data)
	for generation := 0; generation < data.Generations; generation++ {
		iterate(pop, handler, data)
	}
	return pop
}

func createPopulation[T any](handler Handler[T], data *ExperimentData[T]) []Solution[T] {
	result := make([]Solution[T], data.PopulationSize)
	for i := 0; i < data.PopulationSize; i++ {
		result[i] = handler.NewSolution()
	}
	return result
}

func iterate[T any](pop []Solution[T], handler Handler[T], data *ExperimentData[T]) {
	var (
		n                       = len(pop)
		crossoverInChan         = make(chan [2]Solution[T])
		crossoverToMutationChan = make(chan Solution[T])
		out                     = make(chan Solution[T])
	)

	go doCrossover[T](crossoverInChan, crossoverToMutationChan, data.CrossoverProbability, handler.Cross)
	go doMutation[T](crossoverToMutationChan, out, data.MutationProbability, handler.Mutate)

	// feed in
	for _, pair := range data.CrossoverSelectionStrategy(pop, n/2) {
		crossoverInChan <- pair
	}
	close(crossoverInChan)

	// read out
	var popForSelection []Solution[T]
	for sol := range out {
		popForSelection = append(popForSelection, sol)
	}

	copy(pop, data.NextGenerationSelectionStrategy(popForSelection, n))
}

func doCrossover[T any](
	in chan [2]Solution[T],
	out chan Solution[T],
	probability float64,
	crossoverFunc func(s1, s2 Solution[T]) [2]Solution[T],
) {
	for pair := range in {
		if rng.Float64() > probability {
			continue
		}
		resultPair := crossoverFunc(pair[0], pair[1])
		out <- resultPair[0]
		out <- resultPair[1]
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
		if rng.Float64() > probability {
			continue
		}
		out <- mutateFunc(sol)
	}
	close(out)
}
