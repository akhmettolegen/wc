package file_reader

import (
	"bufio"
	"io"
	"os"
)

type FileCounter interface {
	BytesCount() (int, error)
	LinesCount() (int, error)
	WordsCount() (int, error)
	CharsCount() (int, error)
}

type FileReader struct {
	File *os.File
}

func New(file *os.File) *FileReader {
	return &FileReader{File: file}
}

func (f *FileReader) BytesCount() (int, error) {
	bytes, err := io.ReadAll(f.File)
	if err != nil {
		return 0, err
	}

	return len(bytes), err
}

func (f *FileReader) LinesCount() (int, error) {
	var linesCount int
	sc := bufio.NewScanner(f.File)
	for sc.Scan() {
		linesCount++
	}

	return linesCount, nil
}

func (f *FileReader) WordsCount() (int, error) {
	var wordsCount int
	sc := bufio.NewScanner(f.File)
	sc.Split(bufio.ScanWords)
	for sc.Scan() {
		wordsCount++
	}

	return wordsCount, nil
}

func (f *FileReader) CharsCount() (int, error) {
	var charsCount int
	sc := bufio.NewScanner(f.File)
	sc.Split(bufio.ScanRunes)
	for sc.Scan() {
		charsCount++
	}

	return charsCount, nil
}
