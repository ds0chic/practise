package practise

func Reduce[T any, R any](value []T, reducer func(result R, value T, i int) R, initial R) R {
	result := initial
	for i, value := range value {
		result = reducer(result, value, i)
	}
	return result
}
