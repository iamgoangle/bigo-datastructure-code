package main

import (
	"fmt"
)

// Factorial

func main() {
	fmt.Println(factorial(5))
}

func factorial(n int) int {
	f := 1
	for i := 2; i <= n; i++ {
		f = f * i
	}

	return f
}
