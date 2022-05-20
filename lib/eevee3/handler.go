package eevee3

type Crossover[TUnderlying any] interface {
	// Cross takes 2 solutions instances and performs a crossover on them
	Cross(solution1, solution2 Solution[TUnderlying]) (Solution[TUnderlying], Solution[TUnderlying])
}

type Mutation[TUnderlying any] interface {
	// Mutate takes a single solutions instance and performs a mutation on it
	Mutate(solution Solution[TUnderlying]) Solution[TUnderlying]
}

// Handler is the interface of the handler used for operating a problem
type Handler[TUnderlying any] interface {
	Crossover[TUnderlying]
	Mutation[TUnderlying]

	// NewSolution returns a single solutions instance
	NewSolution() Solution[TUnderlying]
}
