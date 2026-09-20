package practise

func find[T any](find func(value T, i int) bool, values []T) (T, bool) {
	var zero T
	for i, value := range values {
		if find(value, i) {
			return value, true
		}
	}
	return zero, false
}
