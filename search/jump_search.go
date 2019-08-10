package main

// https://www.geeksforgeeks.org/jump-search/
// O(√ n)

import (
	"fmt"
	"math"
	"time"
)

/*
	step 1: find len of array / set boundary left = 0 and right = square root of array len
	step 2: compare value of bucket, last item in bucket with `key`, if `key` greather than, jump to next bucket
	step 3: when jump step left = right and right += square root of array len
	step 4: if right (last bucket item) > `key`, begin linear search
*/

func jumpSearch(arr []int, key int) int {
	// O(1)
	size := len(arr) - 1
	left := 0
	right := int(math.Sqrt(float64(size)))

	// O(n)
	for arr[right] < key && right < size {
		left = right
		right += int(math.Sqrt(float64(size)))

		if right > size {
			right = size
		}
	}

	for i := left; i <= right; i++ {
		if arr[i] == key {
			return i
		}
	}

	return -1
}

func main() {
	t := time.Now()
	arr := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610}
	key := 55

	index := jumpSearch(arr, key)
	fmt.Printf("find %v, found at index = %v\n", key, index)
	fmt.Println(time.Since(t))
}
