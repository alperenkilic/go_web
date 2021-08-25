package main

import (
	"fmt"
)

func main() {
	var kare Square
	kare.Kenar = 12
	fmt.Println(SquareAlan(kare))
}

type Square struct {
	Kenar int
}

type Triange struct {
	Kenar int
	High  int
}

func SquareAlan(s Square) int {
	return s.Kenar * s.Kenar
}

func TriangleAlan(t Triange) int {
	return t.Kenar * t.High / 2
}
