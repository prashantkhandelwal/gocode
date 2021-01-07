package main

import "fmt"

func product(n ...int) int {
	p := 1

	for _, i := range n {
		p *= i
	}

	return p
}

func main() {

	// We can also pass a slice in a variadic function.
	n := []int{2, 2, 2}
	fmt.Println(product(n...))
	fmt.Println(product(1, 2, 3))

}
