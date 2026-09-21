package practise

func Count[T comparable](value []T, target T) int {
	count := 0
	for _, v := range value {
		if v == target {
			count++
		}
	}
	return count
}
