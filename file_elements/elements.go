package file_elements

import (
	"bufio"
	"os"
)

const (
	ElementBytes = "bytes"
	ElementRunes = "runes"
	ElementWords = "words"
	ElementLines = "lines"
)

type Counter interface {
	Count() (int, error)
}

type Element struct {
	file    *os.File
	element string
}

func New(file *os.File, element string) *Element {
	return &Element{file, element}
}

// Count counts elements in file
func (e *Element) Count() (int, error) {
	splitterMap := map[string]bufio.SplitFunc{
		ElementBytes: bufio.ScanBytes,
		ElementRunes: bufio.ScanRunes,
		ElementWords: bufio.ScanWords,
		ElementLines: bufio.ScanLines,
	}

	scanner := bufio.NewScanner(e.file)

	var count int
	scanner.Split(splitterMap[e.element])
	for scanner.Scan() {
		count++
	}

	return count, nil
}
