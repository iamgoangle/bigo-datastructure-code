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
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3

	rotate(nums, k)
}

func rotate(nums []int, k int) {
	for i := 0; i < k; i++ {
		prev := nums[len(nums)-1]

		for j := 0; j < len(nums); j++ {
			temp := nums[j]
			nums[j] = prev
			prev = temp
		}
	}

	fmt.Println(nums)
}
