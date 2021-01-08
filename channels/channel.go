package main

import (
	"fmt"
	"math/rand"
	"time"
)

func CalculateValue(values chan int) {
	rand.Seed(time.Now().UnixNano())
	value := rand.Intn(10)
	fmt.Println("Calculate Random Value: ", value)
	values <- value
}

func main() {
	values := make(chan int)
	defer close(values)

	go CalculateValue(values)

	value := <-values
	fmt.Println(value)

}
