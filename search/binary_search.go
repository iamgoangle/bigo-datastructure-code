package main

import (
	"fmt"
	"time"
)

// https://www.geeksforgeeks.org/binary-search/
// O(Log n)

func main() {

	t := time.Now()
	arr := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610}
	key := 55

	index := binarySearch(arr, key)
	fmt.Printf("find %v, found at index = %v\n", key, index)
	fmt.Println(time.Since(t))
}

func binarySearch(arr []int, key int) int {
	size := len(arr)
	left := 0
	right := size - 1

	for left <= size {
		mid := left + (right-left)/2
		if key > arr[mid] {
			left = mid + 1
		} else if key < arr[mid] {
			right = mid - 1
		} else {
			return mid
		}
	}

	return -1
}
