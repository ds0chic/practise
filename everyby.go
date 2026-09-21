package practise

func EveryBy[T any](value []T, predicate func(value T, i int) bool) bool {
	for i, value := range value {
		if !predicate(value, i) {
			return false
		}
	}
	return true
}
