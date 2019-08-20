package main

/*
	Given an array nums, write a function to move all 0's to the end of it while maintaining the relative order of the non-zero elements.

	Example:

	Input: [0,1,0,3,12]
	Output: [1,3,12,0,0]
	Note:

	You must do this in-place without making a copy of the array.
	Minimize the total number of operations.
*/

import (
	"fmt"
)

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
}

func moveZeroes(nums []int) {
	zeroCount := 0
	ans := []int{}

	// find zero
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			zeroCount += 1
		} else {
			ans = append(ans, nums[i])
		}
	}

	// push zero
	for j := zeroCount - 1; j < len(nums)-zeroCount; j++ {
		ans = append(ans, 0)
	}

	fmt.Println(ans)
}
