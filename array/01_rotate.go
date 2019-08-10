package main

import (
	"fmt"
)

// https://www.geeksforgeeks.org/array-rotation/

/*
	temp = [1]
	lists = [2, 3, 4, 5, 6, 7]
	and push `temp` to index `lists[6]``
	so [2, 3, 4, 5, 6, 7, 1]

	temp = [2]
	lists = [3, 4, 5, 6, 7, 1]
	and push `temp` to index `lists[6]`
	so [3, 4, 5, 6, 7, 1, 2]

	time complexity = O(N) * O(N)
*/
func leftRotate(l []int, n int) {
	temp := l[0]
	i := 0
	for i = 0; i < n; i++ {
		l[i] = l[i+1]
	}

	// [2 3 4 5 6 7 7]
	// [3 4 5 6 7 1 1]
	fmt.Println(l)

	// 1st = lists[6] = 1
	// 2nd = lists[6] = 2
	l[i] = temp
}

func rotate(l []int, d, n int) {
	for i := 0; i < d; i++ {
		leftRotate(l, n)
	}
}

func main() {
	lists := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(lists, 2, len(lists)-1)
	fmt.Println("Result => ", lists) // Result =>  [3 4 5 6 7 1 2]
}
