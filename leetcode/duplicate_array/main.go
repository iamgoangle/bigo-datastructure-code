package main

import (
	"fmt"
)

/*
  Input: [1,2,3,1]
  Output: true

  https://leetcode.com/explore/featured/card/top-interview-questions-easy/92/array/578/
*/

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(containsDuplicate(nums))
}

func containsDuplicate(nums []int) bool {
	hTable := make(map[int]int)
	for _, v := range nums {
		hTable[v] = 1
	}

	return len(hTable) == len(nums)
}
