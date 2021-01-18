package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	err := filepath.Walk("/home/prashant/file/1",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				log.Fatal(err)
			}
			if info.IsDir() == false {
				fmt.Println(path, info.Size())
			}
			return nil
		})

	if err != nil {
		log.Fatal(err)
	}
}
