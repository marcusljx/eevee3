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
	// TODO
}
