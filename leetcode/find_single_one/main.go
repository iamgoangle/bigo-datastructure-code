package main

import (
	"fmt"
)

/*
  Given a non-empty array of integers, every element appears twice except for one. Find that single one.

  Note:

  Your algorithm should have a linear runtime complexity. Could you implement it without using extra memory?

  Example 1:

  Input: [2,2,1]
  Output: 1

  https://leetcode.com/explore/featured/card/top-interview-questions-easy/92/array/549/
*/

func main() {
	nums := []int{2, 2, 1, 1, 3}
	fmt.Println(singleNumber(nums))
}

func singleNumber(nums []int) int {
	result := make(map[int]int)
	for _, v := range nums {
		result[v] += 1
	}

	fmt.Println(result)

	for k, v := range result {
		if v == 1 {
			return k
		}
	}

	return -1
}
