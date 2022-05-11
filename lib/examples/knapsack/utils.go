package knapsack

func min[T int | float64](val T, values ...T) T {
	for _, y := range values {
		if y < val {
			val = y
		}
	}
	return val
}

func sum[T int | float64](values ...T) T {
	var result T = 0
	for _, x := range values {
		result += x
	}
	return result
}

func sliceMap[T any, U any](slice []T, mapFunc func(T) U) []U {
	result := make([]U, len(slice))
	for i, x := range slice {
		result[i] = mapFunc(x)
	}
	return result
}

func sliceFold[T any, U any](initialValue U, slice []T, foldFunc func(current U, next T) U) U {
	result := initialValue
	for _, x := range slice {
		result = foldFunc(result, x)
	}
	return result
}
