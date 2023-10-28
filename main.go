package main

import (
	"bufio"
	"fmt"
	"github.com/akhmettolegen/coding-challenges/wc/file_reader"
	"log"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var command, op1, op2 string
	_, _ = fmt.Fscanln(in, &command, &op1, &op2)

	var flag, path string
	if op2 == "" {
		path = op1
	}

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fileReader := file_reader.New(file)

	if command != "ccwc" {
		log.Fatal("invalid command")
	}

	switch flag {
	case "-c":
		bytesCount, err := fileReader.BytesCount()
		if err != nil {
			log.Fatal(err)
		}

		_, _ = fmt.Fprintln(out, bytesCount, path)
	case "-l":
		linesCount, err := fileReader.LinesCount()
		if err != nil {
			log.Fatal(err)
		}

		_, _ = fmt.Fprintln(out, linesCount, path)
	case "-w":
		wordsCount, err := fileReader.WordsCount()
		if err != nil {
			log.Fatal(err)
		}

		_, _ = fmt.Fprintln(out, wordsCount, path)
	case "-m":
		charsCount, err := fileReader.CharsCount()
		if err != nil {
			log.Fatal(err)
		}

		_, _ = fmt.Fprintln(out, charsCount, path)
	case "":
		bytesCount, err := fileReader.BytesCount()
		if err != nil {
			log.Fatal(err)
		}

		linesCount, err := fileReader.LinesCount()
		if err != nil {
			log.Fatal(err)
		}

		wordsCount, err := fileReader.WordsCount()
		if err != nil {
			log.Fatal(err)
		}

		_, _ = fmt.Fprintln(out, bytesCount, linesCount, wordsCount, path)
	default:
		log.Fatal("invalid flag")
	}
}
