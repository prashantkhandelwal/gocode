package main

import (
	"fmt"
	"math/rand"
	"time"
)

func CalculateValue(c chan int) {
	rand.Seed(time.Now().UnixNano())
	value := rand.Intn(10)
	fmt.Println("Calculated Random Value:", value)
	time.Sleep(1000 * time.Millisecond)
	c <- value

	fmt.Println("Only Executes after another goroutine performs a receive on the channel")
}

func main() {
	valueChannel := make(chan int)
	defer close(valueChannel)

	go CalculateValue(valueChannel)
	go CalculateValue(valueChannel)

	values := <-valueChannel
	fmt.Println(values)
}
