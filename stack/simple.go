package main

import (
	"fmt"
)

func main() {
	stack := []string{}

	stack = append(stack, "a")
	stack = append(stack, "b")
	stack = append(stack, "c")
	stack = append(stack, "d")
	stack = append(stack, "e")

	fmt.Println(stack)

	// pop
	fmt.Println("value", stack[len(stack)-1]) // pop
	stack = stack[:len(stack)-1]
	fmt.Println("remain", stack)
}
