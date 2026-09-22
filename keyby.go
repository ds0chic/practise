package practise

func KeyBy[T any, K comparable](values []T, key func(value T) K) map[K]T {
	result := make(map[K]T)
	for _, value := range values {
		k := key(value)
		result[k] = value
	}
	return result
}
