package main

import "fmt"

type filter func(int) bool

func isEven(n int) bool {
	return n%2 == 0
}

func isOdd(n int) bool {
	return n%2 == 1
}

func main() {
	var value filter

	value = isEven

	fmt.Println(value(4))
}
