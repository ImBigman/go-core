// Package fibo формирует массив чисел и выводит значения по порядковому номеру.
package fibo

func Numbers(x int, y int) int {
	a := make([]int, x, x)
	a[0], a[1] = 0, 1

	for i := 2; i < x; i++ {
		a[i] = a[i-1] + a[i-2]
	}

	return a[y-1]
}
