package eevee3

func MutateRandomIndex[T any](slice []T, changeFunc func(T) T) {
	targetIdx := rng.Int() % len(slice)
	slice[targetIdx] = changeFunc(slice[targetIdx])
}
