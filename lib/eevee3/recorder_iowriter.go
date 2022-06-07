package eevee3

import "io"

// NewIOWriterRecorder creates a new IOWriterRecorder
func NewIOWriterRecorder[T any](writer io.Writer, config *RecorderConfig) *IOWriterRecorder[T] {
	return &IOWriterRecorder[T]{
		w:       writer,
		include: *config,
	}
}

type IOWriterRecorder[T any] struct {
	w       io.Writer
	include RecorderConfig
}

func (r *IOWriterRecorder[T]) ExperimentStart(controller Controller[T]) {
	if !r.include.ExperimentStart {
		return
	}
	// TODO
}

func (r *IOWriterRecorder[T]) InitialPopulation(population []Solution[T]) {
	if !r.include.InitialPopulation {
		return
	}
	// TODO
}

func (r *IOWriterRecorder[T]) Generation(generation int) {
	if !r.include.Generation {
		return
	}
	// TODO
}

func (r *IOWriterRecorder[T]) Crossover(in1, in2, out1, out2 Solution[T]) {
	if !r.include.Crossover {
		return
	}
	// TODO
}

func (r *IOWriterRecorder[T]) Mutate(in, out Solution[T]) {
	if !r.include.Mutate {
		return
	}
	// TODO
}

func (r *IOWriterRecorder[T]) Selection(before, after []Solution[T]) {
	if !r.include.Selection {
		return
	}
	// TODO
}

func (r *IOWriterRecorder[T]) Terminate(result Result[T]) {
	if !r.include.Terminate {
		return
	}
	// TODO
}
