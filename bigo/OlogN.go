package main

import (
	"errors"
	"fmt"
)

/*
	lists := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	find := 3

	left = 0
	mid = 4
	right = 9

	left = 0
	mid = 1
	right = 3

	left = 2
	mid = 2
	right = 3
*/

func main() {
	lists := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	find := 7

	fmt.Println(binarySearch(lists, find))
}

func binarySearch(lists []int, find int) (string, error) {
	low := 0
	high := len(lists) - 1

	for {
		if low <= high {
			var mid int = low + (high-low)/2

			if find < lists[mid] {
				high = mid - 1
			} else if find > lists[mid] {
				low = mid + 1
			} else {
				return "found", nil
			}

			fmt.Println("low => ", low)
			fmt.Println("mid => ", mid)
			fmt.Println("high => ", high)
			fmt.Println("====")
		} else {
			return "", errors.New("not found")
		}
	}
}
