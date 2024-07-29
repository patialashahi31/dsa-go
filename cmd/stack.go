package main

import (
	"fmt"

	"github.com/patialashahi31/dsa-go/stackqueue"
)

func main() {
	stack := stackqueue.Stack{}
	stack.Push(1)
	fmt.Println(stack.Top())
	stack.Push(2)
	fmt.Println(stack.Top())
	stack.Push(3)
	fmt.Println(stack.Top())
	stack.Pop()
	fmt.Println(stack.Top())
}
