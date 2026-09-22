package practise

func UniqBy[T any, K comparable](value []T, key func(value T) K) []T {
	result := make([]T, 0)
	seen := make(map[K]bool)
	for _, value := range value {
		k := key(value)
		if !seen[k] {
			seen[k] = true
			result = append(result, value)
		}
	}
	return result
}
