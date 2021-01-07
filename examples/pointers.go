package main

import (
	"fmt"
)

func withPointer(x *int) {
	*x = *x * *x
}

type complex struct {
	x, y int
}

func newComplexType(x, y int) *complex {
	return &complex{x, y}
}

// To read the value from the struct, you need to
// define a function which can return the value.
func (nc *complex) GetX() int {
	return nc.x
}

func main() {
	x := 2
	withPointer(&x)
	fmt.Println(x)

	w := newComplexType(4, -5)
	fmt.Println(*w)
	fmt.Println(w)
	fmt.Println(w.GetX())
}
