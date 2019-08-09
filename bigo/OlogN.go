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
	find := 3

	fmt.Println(binarySearch(lists, 0, len(lists)-1, find))
}

func binarySearch(lists []int, left, right, find int) (string, error) {
	if right < left {
		return "", errors.New("left must not greather than right")
	}

	mid := left + ((right - left) / 2)
	fmt.Println("left = ", left)
	fmt.Println("mid = ", mid)
	fmt.Println("right = ", right)
	fmt.Println("=====")

	if lists[mid] == find {
		return fmt.Sprintf("found %v", lists[mid]), nil
	}

	if lists[mid] > find {
		return binarySearch(lists, left, mid-1, find)
	}

	if lists[mid] < find {
		return binarySearch(lists, mid+1, right, find)
	}

	return "", nil
}
