package file_content

import (
	"errors"
	"strings"
)

const (
	ElementBytes = "bytes"
	ElementRunes = "runes"
	ElementWords = "words"
	ElementLines = "lines"
)

type Counter interface {
	CountElements(element string) (int, error)
}

type Content struct {
	content []byte
}

func New(content []byte) *Content {
	return &Content{content}
}

// CountElements counts elements in byte slice
func (e *Content) CountElements(element string) (int, error) {
	switch element {
	case ElementBytes:
		return countBytes(e.content), nil
	case ElementRunes:
		return countRunes(e.content), nil
	case ElementWords:
		return countWords(e.content), nil
	case ElementLines:
		return countLines(e.content), nil
	default:
		return 0, errors.New("invalid element")
	}
}

func countBytes(content []byte) int {
	return len(content)
}

func countRunes(content []byte) int {
	return len([]rune(string(content)))
}

func countWords(content []byte) int {
	return len(strings.Fields(string(content)))
}

func countLines(content []byte) int {
	var count int
	for _, c := range string(content) {
		if c == '\n' {
			count++
		}
	}
	return count
}
