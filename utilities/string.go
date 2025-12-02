package utilities

// ChunkString breaks the given string into chunks of equal size. If it can't break into equal-size chunks, it returns nil.
func ChunkString(s string, chunkSize int) []string {
	if len(s)%chunkSize != 0 {
		return nil
	}

	var chunks []string

	for i := 0; i < len(s); i += chunkSize {
		end := min(i+chunkSize, len(s))
		chunks = append(chunks, s[i:end])
	}

	return chunks
}
