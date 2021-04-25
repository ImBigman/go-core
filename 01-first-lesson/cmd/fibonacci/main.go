package main

import (
	"fmt"
	"go-core/01-first-lesson/pkg/fibo"
)

func main() {
	var d, e = 8, 19

	printFiboNUm(d)
	printFiboNUm(e)
	printFiboNUm(-1)
	printFiboNUm(22)
}

// printFiboNUm - выводит на экран результат или предупреждения
// с - порядковый номер
func printFiboNUm(c int) {
	if 0 < c && c <= 20 {
		fmt.Println(fibo.Num(c))
	} else {
		fmt.Println("Недопустимое значение")
	}
}
