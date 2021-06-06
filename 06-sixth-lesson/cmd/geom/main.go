package main

import (
	"fmt"
	"go-core/06-sixth-lesson/pkg/measuring"
)

func main() {
	coord := measuring.Coord{X1: 1.23, X2: 3.45, Y1: 6.78, Y2: 9.1}
	if coord.X1 < 0 || coord.X2 < 0 || coord.Y1 < 0 || coord.Y2 < 0 {
		fmt.Println("Координаты не могут быть меньше нуля")
	}
	fmt.Printf("Расстояние между точками: %f", measuring.Distance(coord))
}
