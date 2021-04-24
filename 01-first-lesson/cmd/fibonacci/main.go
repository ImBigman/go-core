package main

import (
	"fmt"
	"go-core/01-first-lesson/pkg/fibo"
)

func main() {
	var d, e = 8, 19

	print(d)
	print(e)
	print(-1)
	print(22)
}

// print - выводит на экран результат или предупреждения
// с - порядковый номер
func print(c int) {

	if 0 < c && c <= 20 {
		fmt.Println(fibo.Num(c))
	} else {
		fmt.Println("Недопустимое значение")
	}

}
