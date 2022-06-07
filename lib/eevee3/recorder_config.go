package eevee3

var (
	RecordEverything = &RecorderConfig{
		ExperimentStart:   true,
		InitialPopulation: true,
		Generation:        true,
		Crossover:         true,
		Mutate:            true,
		Selection:         true,
		Terminate:         true,
	}
)

type RecorderConfig struct {
	ExperimentStart   bool
	InitialPopulation bool
	Generation        bool
	Crossover         bool
	Mutate            bool
	Selection         bool
	Terminate         bool
}
