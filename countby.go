package practise

func CountBy[T any](values []T, predicate func(value T, i int) bool) int {
	count := 0
	for i, value := range values {
		if predicate(value, i) {
			count++
		}
	}
	return count
}
