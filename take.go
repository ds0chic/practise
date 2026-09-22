package practise

func Take[T any](values []T, count int) []T {
	if count <= 0 {
		return nil
	}
	if count >= len(values) {
		return values
	}
	return values[:count]
}
