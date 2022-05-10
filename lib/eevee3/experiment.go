package eevee3

// ExperimentData represents inner structure data about the experiment
type ExperimentData[T any] struct {
	// Generations is the number of generation cycles
	// to run the experiment for
	Generations int

	// PopulationSize represents the number of solutions
	// in each generation corpus
	PopulationSize int

	// MutationProbability is the probability that
	// each solution has to mutate on each generation cycle
	MutationProbability float64

	// CrossoverSelectionStrategy is the strategy used to select pairs
	// for crossover
	CrossoverSelectionStrategy PairwiseSelectionStrategy[T]

	// CrossoverProbability is the probability that
	// each pair of solutions has to crossover together
	CrossoverProbability float64

	// NextGenerationSelectionStrategy is the strategy used in each iteration
	// to select the next corpus of solutions
	NextGenerationSelectionStrategy SubgroupSelectionStrategy[T]
}