package main

import (
	"fmt"
	"go-core/04-fourth-lesson/pkg/list"
)

func main() {

	l := list.New()
	l.Push(list.Elem{Val: 3})
	l.Push(list.Elem{Val: 2})
	l.Push(list.Elem{Val: 1})
	l = l.Reverse()
	fmt.Println(l)
}

type List list.List
