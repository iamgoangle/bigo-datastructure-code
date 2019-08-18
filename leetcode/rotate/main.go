package main

import (
	"fmt"
)

/*
  Input: [1,2,3,4,5,6,7] and k = 3
  Output: [5,6,7,1,2,3,4]
  Explanation:
  rotate 1 steps to the right: [7,1,2,3,4,5,6]
  rotate 2 steps to the right: [6,7,1,2,3,4,5]
  rotate 3 steps to the right: [5,6,7,1,2,3,4]

  https://leetcode.com/explore/featured/card/top-interview-questions-easy/92/array/646/
*/

func main() {
	// nums := []int{1,2,3,4,5,6,7}
	nums := []int{-1, -100, 3, 99}

	// k := 3
	k := 2

	rotate(nums, k)
}

func rotate(nums []int, k int) {
	slice := []int{}
	for i := len(nums) - k; i < len(nums); i++ {
		slice = append(slice, nums[i])
	}

	shift := []int{}
	for j := 0; j < len(nums)-k; j++ {
		shift = append(shift, nums[j])
	}

	result := append(slice, shift...)
	fmt.Println(result)
}
