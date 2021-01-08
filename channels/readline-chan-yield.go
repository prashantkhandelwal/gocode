package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func ReadLines(path string) (<-chan string, error) {
	fobj, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(fobj)
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	chnl := make(chan string)
	go func() {
		for scanner.Scan() {
			chnl <- scanner.Text()
		}
		close(chnl)
	}()

	return chnl, nil
}

func main() {
	reader, err := ReadLines("README.md")
	if err != nil {
		log.Fatal(err)
	}

	for l := range reader {
		fmt.Println(l)
		break
	}
}
