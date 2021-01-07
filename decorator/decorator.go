package main

import "fmt"

func hello() {
	fmt.Println("Welcome")
}

func coolFunc(a func()) {
	a()
}

func main() {
	fmt.Printf("Type: %T\n", hello)

	coolFunc(hello)
}
