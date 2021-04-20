package main

import "go-core/01-first-lesson/pkg/fibo"

func main() {
	var a int = 8
	var b int = 19

	fibo.FiboNumbers(a)
	fibo.FiboNumbers(b)
	fibo.FiboNumbers(21)
	fibo.FiboNumbers(-1)
}
