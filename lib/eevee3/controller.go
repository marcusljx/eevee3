package eevee3

// Controller represents inner structure data about the experiment,
// as well as strategies for handling experiment execution.
type Controller[T any] struct {
	Recorder Recorder[T]

	// GenerationCycles is the number of generation cycles
	// to run the experiment for
	GenerationCycles int

	// PopulationSize represents the number of solutions
	// in each generation corpus
	PopulationSize int

	// MutationProbability is the probability that
	// each solutions has to mutate on each generation cycle
	MutationProbability float64

	// CrossoverProbability is the probability that
	// each pair of solutions has to crossover together
	CrossoverProbability float64

	// CrossoverSelectionStrategy is the strategy used to select pairs
	// for crossover
	CrossoverSelectionStrategy PairwiseSelectionStrategy[T]

	// NextGenerationSelectionStrategy is the strategy used in each iteration
	// to select the next corpus of solutions
	NextGenerationSelectionStrategy SubgroupSelectionStrategy[T]

	// TerminationConditions allows for input of conditions
	// that may allow for early termination of the experiment
	TerminationConditions []PopulationPredicate[T]
}
