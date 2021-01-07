// We have a type A, with a field year and a method Greet. We have a second type, B which embeds an A,
// thus callers see B‘s methods overlaid on A‘s because A is embedded, as a field, within B,
// and B can provide its own Greet method, obscuring that of A.

// But embedding isn’t just for methods, it also provides access to an embedded type’s fields.
// As you see, because both A and B are defined in the same package,
// B can access A‘s private year field as if it were declared inside B.

package main

import (
	"fmt"
)

type A struct {
	year int
}

func (a A) Greet() {
	fmt.Println("Hello Golang", a.year)
}

type B struct {
	A
}

func (b B) Greet() {
	fmt.Println("Welcome to the world of programming", b.year)
}

func main() {
	var a A
	var b B

	a.year = 2019
	b.year = 2020

	a.Greet()
	b.Greet()
}
