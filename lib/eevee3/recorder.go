package eevee3

// Recorder is a logger interface for visibility into the iteration process.
// All method implementations must be concurrent-safe.
type Recorder[T any] interface {
	// ExperimentStart marks the start of the experiment
	ExperimentStart(controller Controller[T])

	// InitialPopulation records the initial population
	InitialPopulation(population []Solution[T])

	// Generation records whenever a new generation cycle begins
	Generation(generation int)

	// Crossover records each crossover event
	Crossover(in1, in2, out1, out2 Solution[T])

	// Mutate records each mutation event
	Mutate(in, out Solution[T])

	// Selection records the selection process at the end of each cycle
	Selection(before, after []Solution[T])

	// Terminate records a termination
	Terminate(result Result[T])
}
