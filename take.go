package practise

func Take[T any](values []T, count int) []T {
	if count < 0 {
		panic("practise.Take: count must not be negative")
	}
	if count == 0 || len(values) == 0 {
		return make([]T, 0)
	}
	if count >= len(values) {
		result := make([]T, len(values))
		copy(result, values)
		return result
	}
	result := make([]T, count)
	copy(result, values)
	return result
}
