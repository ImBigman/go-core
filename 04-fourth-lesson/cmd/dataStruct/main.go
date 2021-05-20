package main

import (
	"fmt"
	"go-core/04-fourth-lesson/pkg/list"
)

func main() {

	l := list.New()
	l.Push(list.Elem{Val: 4})
	l.Push(list.Elem{Val: 3})
	l.Push(list.Elem{Val: 2})
	l.Push(list.Elem{Val: 1})
	fmt.Println("Before: ", l)
	fmt.Println("After: ", l.Reverse())
}

type List list.List
