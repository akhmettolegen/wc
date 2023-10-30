package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/akhmettolegen/coding-challenges/wc/file_content"
	"io"
	"log"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	flag, path, err := parseCmd(in)
	if err != nil {
		log.Fatalf("error while parsing command: %v", err)
	}

	fileBytes, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("error while reading file: %v", err)
	}

	var content file_content.Counter
	content = file_content.New(fileBytes)

	switch flag {
	case "-c":
		count, err := content.CountElements(file_content.ElementBytes)
		if err != nil {
			log.Fatalf("error while count: %v", err)
		}

		_, _ = fmt.Fprintln(out, count, path)
	case "-l":
		count, err := content.CountElements(file_content.ElementLines)
		if err != nil {
			log.Fatalf("error while count: %v", err)
		}

		_, _ = fmt.Fprintln(out, count, path)
	case "-w":
		count, err := content.CountElements(file_content.ElementWords)
		if err != nil {
			log.Fatalf("error while count: %v", err)
		}

		_, _ = fmt.Fprintln(out, count, path)
	case "-m":
		count, err := content.CountElements(file_content.ElementRunes)
		if err != nil {
			log.Fatalf("error while count: %v", err)
		}

		_, _ = fmt.Fprintln(out, count, path)
	case "all":
		bytesC, err := content.CountElements(file_content.ElementBytes)
		if err != nil {
			log.Fatalf("error while count %s: %v", file_content.ElementBytes, err)
		}

		linesC, err := content.CountElements(file_content.ElementLines)
		if err != nil {
			log.Fatalf("error while count %s: %v", file_content.ElementLines, err)
		}

		wordsC, err := content.CountElements(file_content.ElementWords)
		if err != nil {
			log.Fatalf("error while count %s: %v", file_content.ElementWords, err)
		}

		_, _ = fmt.Fprintln(out, bytesC, linesC, wordsC, path)
	default:
		log.Fatal("invalid flag")
	}
}

func parseCmd(in io.Reader) (string, string, error) {
	var command, op1, op2, op3, op4 string
	_, _ = fmt.Fscanln(in, &command, &op1, &op2, &op3, &op4)

	var flag, path string
	if command == ">ccwc" {
		flag, path = op1, op2
		if op2 == "" {
			path = op1
			flag = "all"
		}
	} else if command == ">cat" {
		path = op1
		flag = op4
	} else {
		return "", "", errors.New("invalid command")
	}

	return flag, path, nil
}
