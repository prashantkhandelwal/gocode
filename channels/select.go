package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(5 * time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "two"
	}()

	// We’ll use select to await both of these values simultaneously,
	// printing each one as it arrives.
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("Recevied", msg1)
		case msg2 := <-c2:
			fmt.Println("Recevied", msg2)
		}
	}
}
