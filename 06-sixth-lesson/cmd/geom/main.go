package main

import (
	"fmt"
	"go-core/06-sixth-lesson/pkg/calc"
)

func main() {
	coord := calc.Coord{X1: 1.23, X2: 3.45, Y1: 6.78, Y2: 9.1}
	res := calc.Distance(coord)
	if res == -1 {
		fmt.Println("Координаты не могут быть меньше нуля")
	}
	fmt.Printf("Расстояние между точками: %f", res)
}
