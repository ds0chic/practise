package practise

func Filter[T any](filter func(vlaue T, i int) bool, values []T) []T {
	var result []T
	for i, value := range values {
		if filter(value, i) {
			result = append(result, value)
		}
	}
	return result
}
