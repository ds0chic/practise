package practise

func Contains[T comparable](list []T, item T) bool {
	for _, v := range list {
		if v == item {
			return true
		}
	}
	return false
}
