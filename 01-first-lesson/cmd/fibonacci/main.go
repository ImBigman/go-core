package main

import (
	"fmt"
	"go-core/01-first-lesson/pkg/fibo"
)

func main() {
	var d, e = 8, 19

	print(40, d)
	print(20, e)
	print(20, -1)
	print(20, 22)
}

// print - выводит на экран результат или предупреждения
// b - длина массива, с - порядковый номер
func print(b int, c int) {

	if 0 < c && c <= b {
		fmt.Println(fibo.Num(c))
	} else {
		fmt.Println("Недопустимое значение")
	}

}
