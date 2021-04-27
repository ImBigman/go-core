package main

import (
	"fmt"
	"go-core/01-first-lesson/pkg/fibo"
)

func main() {
	var nums = []int{8, 19, -1, 22}
	for _, n := range nums {
		if 0 < n && n <= 20 {
			fmt.Println(fibo.Num(n))
		} else {
			fmt.Println("Недопустимое значение")
		}
	}
}
