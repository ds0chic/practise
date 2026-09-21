package practise

func Flatmap[T any, R any](value []T, mapper func(value T, i int) []R) []R {
	var result []R
	for i, value := range value {
		newvalue := mapper(value, i)
		result = append(result, newvalue...)
	}
	return result
}
