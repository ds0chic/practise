package practise

func Flatten[T any](chunks [][]T) []T {
	var flattened []T
	for _, chunk := range chunks {
		flattened = append(flattened, chunk...)
	}
	return flattened
}
