package main

import "fmt"

// defer make it use as LIFO

func a1() {
	for i := 0; i < 3; i++ {
		defer fmt.Print(i, " ")
	}
}

func a2() {
	for i := 0; i < 3; i++ {
		defer func() {
			fmt.Print(i, " ")
		}()
	}
}

func a3() {
	for i := 0; i < 3; i++ {
		// The anonymous function is not attached to the variable 'i'
		// So we have to pass it at the end to use it's value.
		defer func(n int) {
			fmt.Print(n, " ")
		}(i) // This acts as an input to the defer anon function.
	}
}

func main() {
	a1()
	a2()
	a3()
}
