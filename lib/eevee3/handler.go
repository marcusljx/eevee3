package eevee3

// Handler is the interface of the handler used for operating a problem
type Handler[TUnderlying any] interface {
	// NewSolution returns a single solution instance
	NewSolution() Solution[TUnderlying]

	// Cross takes 2 solution instances and performs a crossover on them
	Cross(solution1, solution2 Solution[TUnderlying]) (Solution[TUnderlying], Solution[TUnderlying])

	// Mutate takes a single solution instance and performs a mutation on it
	Mutate(solution Solution[TUnderlying]) Solution[TUnderlying]
}
