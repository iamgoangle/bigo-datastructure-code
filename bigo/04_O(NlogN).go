package main

import (
	"fmt"
)

// Linearithmic
// 2 Loop
// first loop run as linear from input length
// 2nd loop, cut input that we don't need

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println("i => ", i)
		for j := 1; j < 10; j *= 2 {
			fmt.Println("j => ", j)
		}
	}
}
