package main

import (
	"fmt"
)

/*
	Time complexity : O(n^2)
	Space complexity : O(1)
*/
// To execute Go code, please declare a func main() in a package "main"

func main() {
	nums := []int{2, 11, 7, 15}
	target := 9

	fmt.Println(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
	resp := []int{}

	for i := 0; i < len(nums); i++ {
		a := target - nums[i]
		for j := i + 1; j < len(nums); j++ {
			if nums[j] == a {
				resp = append(resp, i)
				resp = append(resp, j)

				return resp
			}
		}
	}

	return resp
}
