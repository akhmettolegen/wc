package main

import (
	"bufio"
	"fmt"
	"github.com/akhmettolegen/coding-challenges/wc/file_elements"
	"log"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

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
		log.Fatal("invalid command")
	}

	switch flag {
	case "-c":
		file, err := os.Open(path)
		if err != nil {
			log.Fatalf("error while opening file: %v", err)
		}
		defer file.Close()

		elements := file_elements.New(file, file_elements.ElementBytes)
		count, err := elements.Count()
		if err != nil {
			log.Fatalf("error while count %s: %v", file_elements.ElementBytes, err)
		}

		_, _ = fmt.Fprintln(out, count, path)
	case "-l":
		file, err := os.Open(path)
		if err != nil {
			log.Fatalf("error while opening file: %v", err)
		}
		defer file.Close()

		elements := file_elements.New(file, file_elements.ElementLines)
		count, err := elements.Count()
		if err != nil {
			log.Fatalf("error while count %s: %v", file_elements.ElementLines, err)
		}

		_, _ = fmt.Fprintln(out, count, path)
	case "-w":
		file, err := os.Open(path)
		if err != nil {
			log.Fatalf("error while opening file: %v", err)
		}
		defer file.Close()

		elements := file_elements.New(file, file_elements.ElementWords)
		count, err := elements.Count()
		if err != nil {
			log.Fatalf("error while count %s: %v", file_elements.ElementWords, err)
		}

		_, _ = fmt.Fprintln(out, count, path)
	case "-m":
		file, err := os.Open(path)
		if err != nil {
			log.Fatalf("error while opening file: %v", err)
		}
		defer file.Close()

		elements := file_elements.New(file, file_elements.ElementRunes)
		count, err := elements.Count()
		if err != nil {
			log.Fatalf("error while count %s: %v", file_elements.ElementRunes, err)
		}

		_, _ = fmt.Fprintln(out, count, path)
	case "all":
		file, err := os.Open(path)
		if err != nil {
			log.Fatalf("error while opening file: %v", err)
		}
		defer file.Close()

		elements := file_elements.New(file, file_elements.ElementBytes)
		bytesC, err := elements.Count()
		if err != nil {
			log.Fatalf("error while count %s: %v", file_elements.ElementBytes, err)
		}

		file, err = os.Open(path)
		if err != nil {
			log.Fatalf("error while opening file: %v", err)
		}
		defer file.Close()

		elements = file_elements.New(file, file_elements.ElementLines)
		linesC, err := elements.Count()
		if err != nil {
			log.Fatalf("error while count %s: %v", file_elements.ElementLines, err)
		}

		file, err = os.Open(path)
		if err != nil {
			log.Fatalf("error while opening file: %v", err)
		}
		defer file.Close()

		elements = file_elements.New(file, file_elements.ElementWords)
		wordsC, err := elements.Count()
		if err != nil {
			log.Fatalf("error while count %s: %v", file_elements.ElementWords, err)
		}

		_, _ = fmt.Fprintln(out, bytesC, linesC, wordsC, path)
	default:
		log.Fatal("invalid flag")
	}
}
