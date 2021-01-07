package main

import "fmt"

// Here we are naming the return types of the function so it is
// easier to get the sense of the variables are the returning.
func MinMax(x, y int) (min, max int) {
	if x > y {
		min = y
		max = x
	} else {
		min = x
		max = y
	}

	return
}

func main() {
	x := 3
	y := 4

	min, max := MinMax(x, y)
	fmt.Printf("Min: %v and Max: %v\n", min, max)

}
