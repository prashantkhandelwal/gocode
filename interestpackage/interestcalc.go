package main

import (
	"fmt"

	"github.com/prashantkhandelwal/gocode/interestpackage/calc"
)

func main() {
	fmt.Println("Simple Interest Calculator")
	p := 5000.0
	r := 10.0
	t := 1.0
	si := calc.Calculate(p, r, t)
	fmt.Println("Simple interest is: ", si)
}
