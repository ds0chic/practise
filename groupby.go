package practise

func GroupBy[T any, K comparable](value []T, key func(value T) K) map[K][]T {
	result := make(map[K][]T)
	for _, value := range value {
		k := key(value)
		result[k] = append(result[k], value)
	}
	return result
}
