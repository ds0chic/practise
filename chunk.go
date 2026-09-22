package practise

func Chunk[T any](chunk []T, chunkSize int) [][]T {
	if chunkSize <= 0 {
		panic("practise.Chunk: chunkSize must be greater than 0")
	}
	var chunks [][]T
	for i := 0; i < len(chunk); i += chunkSize {
		end := chunkSize + i
		if end > len(chunk) {
			end = len(chunk)
		}
		chunks = append(chunks, chunk[i:end])
	}
	return chunks
}
