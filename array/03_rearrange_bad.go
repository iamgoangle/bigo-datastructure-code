package main

import (
	"fmt"
)

// Input : arr = {-1, -1, 6, 1, 9, 3, 2, -1, 4, -1}
// Output : [-1, 1, 2, 3, 4, -1, 6, -1, -1, 9]

func main() {
	lists := []int{-1, -1, 6, 1, 9, 3, 2, -1, 4, -1}
	size := len(lists)

	reArrange(lists, size)
}

func reArrange(l []int, s int) {
	result := []int{}
	for i := 0; i < s; i++ {
		result = append(result, -1)
	}

	for j := 0; j < s; j++ {
		for k := 1; k < s; k++ {
			if j == l[k] {
				result[j] = j
				break
			}
		}
	}

	fmt.Println(result)
}
