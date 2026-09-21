package practise

func Foreach[T any](value []T, each func(value T, i int)) {
	for i, value := range value {
		each(value, i)
	}
}
