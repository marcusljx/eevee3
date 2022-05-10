package eevee3

// orderedSlice returns a slice of
// incrementing values (starting from 0)
func orderedSlice(n int) []int {
	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = i
	}
	return result
}
