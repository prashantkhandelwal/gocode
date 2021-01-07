package main

import (
	"fmt"
	"sync"
)

func gen(wg *sync.WaitGroup) {
	fmt.Println("Inside goroutine!")
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	go gen(&wg)
	wg.Wait()

	fmt.Println("Finished!")

}
