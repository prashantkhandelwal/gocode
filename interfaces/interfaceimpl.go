package main

import "fmt"

type Guitarist interface {
	PlayGuitar()
}

type BaseGuitarist struct {
	Name string
}

type AcousticGuitarist struct {
	Name string
}

func (b BaseGuitarist) PlayGuitar() {
	fmt.Printf("%s plays the Bass Guitar.\n", b.Name)
}

func (a AcousticGuitarist) PlayGuitar() {
	fmt.Printf("%s plays the Acoustic Guitar.\n", a.Name)
}

func main() {
	playerA := BaseGuitarist{
		Name: "Prashant",
	}

	playerA.PlayGuitar()

	playerB := AcousticGuitarist{
		Name: "Swami",
	}

	playerB.PlayGuitar()

}
