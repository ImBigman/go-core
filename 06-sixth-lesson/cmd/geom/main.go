package main

import (
	"fmt"
	"go-core/06-sixth-lesson/pkg/tapemeasure"
)

func main() {
	coord := tapemeasure.Coord{X1: 1.23, X2: 3.45, Y1: 6.78, Y2: 9.1}
	if coord.X1 < 0 || coord.X2 < 0 || coord.Y1 < 0 || coord.Y2 < 0 {
		fmt.Println("Координаты не могут быть меньше нуля")
	}
	fmt.Printf("Расстояние между точками: %f", tapemeasure.Distance(coord))
}
