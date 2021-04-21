package main

import (
	"fmt"
	"go-core/01-first-lesson/pkg/fibo"
)

func main() {
	var d, e = 8, 19

	printer(40, d)
	printer(20, e)
	printer(20, -1)
	printer(20, 22)
}

// printer - выводит на экран результат или предупреждения
// b - длина массива, с - порядковый номер
func printer(b int, c int) {

	if 0 < c && c <= b {
		fmt.Println(fibo.Numbers(b, c))
	} else {
		fmt.Println("Недопустимое значение")
	}

}
