// Package fibo формирует массив чисел и выводит значения по индексу.
package fibo

import "fmt"

func FiboNumbers(n int) {
	var fiboArray [20]int = [20]int{0, 1}

	for i := 2; i < 20; i++ {
		fiboArray[i] = fiboArray[i-1] + fiboArray[i-2]
	}

	printFibo(n, fiboArray)

}

func printFibo(x int, y [20]int) {
	if 0 < x && x <= 20 {
		fmt.Println(y[x-1])
	} else {
		fmt.Println("Недопустимое значение")
	}

}
