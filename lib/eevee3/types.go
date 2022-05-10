package eevee3

type Scoreable interface {
	Score() float64
}

type Solution[T any] interface {
	Scoreable
	OriginalValue() T
}
