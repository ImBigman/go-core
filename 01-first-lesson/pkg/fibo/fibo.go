// Package fibo возвращает значения по порядковому номеру.
package fibo

// f - порядковый номер искомого значения из ряда
func Num(f int) int {
	b, c := 0, 1
	for i := 0; i <= f-2; i++ {
		b, c = c, b+c
	}

	return b
}
