// When you see a function or a method that expects an empty interface,
// then you can typically pass anything into this function/method.

package main

import "fmt"

func func1(a interface{}) {
	fmt.Println(a)
}

func main() {
	age := 33

	func1(age)
}
