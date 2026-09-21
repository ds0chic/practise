package practise

func FilterMap[T any, R any](value []T, mapper func(value T, i int) (R, bool)) []R {
	var result []R
	for i, value := range value {
		newvalue, ok := mapper(value, i)
		if ok {
			result = append(result, newvalue)
		}
	}
	return result
}
