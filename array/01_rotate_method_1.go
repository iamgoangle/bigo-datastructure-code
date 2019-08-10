package main

import (
	"fmt"
)

// Time complexity : O(n)

func main() {
	lists := []int{1, 2, 3, 4, 5, 6, 7}
	shiftArray(lists, 2, len(lists)-1)
}

func shiftArray(l []int, d, n int) {
	// temp = [1, 2]
	var temp [2]int
	for i := 0; i < d; i++ {
		temp[i] = l[i]
	}

	// Shift rest of the `lists`
	// [3, 4, 5, 6, 7, 6, 7]
	for k := 0; k < n-1; k++ {
		l[k] = l[k+2]
	}

	// store back
	// [3, 4, 5, 6, 7, 1, 2]
	t := 1
	for v := n; v > n-2; v-- {
		l[v] = temp[t]
		t--
	}

	fmt.Println(l)
}
