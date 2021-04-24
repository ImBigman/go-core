package main

import (
	"fmt"
	"go-core/01-first-lesson/pkg/fibo"
)

func main() {
	var d, e = 8, 19

	resultFor(d)
	resultFor(e)
	resultFor(-1)
	resultFor(22)
}

// resultFor - выводит на экран результат или предупреждения
// с - порядковый номер
func resultFor(c int) {
	if 0 < c && c <= 20 {
		fmt.Println(fibo.Num(c))
	} else {
		fmt.Println("Недопустимое значение")
	}
}
