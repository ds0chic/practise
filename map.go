package practise

func Map[T any, R any](values []T, mapper func(value T, i int) R) []R {
	var result []R
	for i, value := range values {
		result = append(result, mapper(value, i))
	}
	return result
}
