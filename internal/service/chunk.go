package service

import (
	"errors"
	"fmt"
)

type ChunkError struct {
	Err error
}

func (c ChunkError) Error() string {
	return fmt.Sprintf("chunk error: %v", c.Err)
}

var (
	chunkSize = 512
)

// Chunk the text to the slice , each slice has length of chunkSize many letters.
func Chunk(text string) ([]string, error) {
	var chunks []string

	if len(text) < 0 {
		return nil, ChunkError{Err: errors.New("empty text")}
	}
	for i := 0; i < len(text); i += chunkSize {
		end := i + chunkSize
		if end > len(text) {
			end = len(text) // Ensure we don't go out of bounds
		}
		chunks = append(chunks, text[i:end])
	}
	return chunks, nil

}
