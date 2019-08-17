package main

/*
	Given an array of integers, return indices of the two numbers such that they add up to a specific target.
	You may assume that each input would have exactly one solution, and you may not use the same element twice.

	Example:

	Given nums = [2, 7, 11, 15], target = 9,
	Because nums[0] + nums[1] = 2 + 7 = 9,
	return [0, 1].
*/

import "fmt"

func main() {
	nums := []int{2, 11, 8, 15, 7}
	target := 9

	fmt.Println(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
	resp := []int{}

	htable := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		htable[nums[i]] = i
	}

	for hKey, arrIdx := range htable {
		myValue := target - hKey
		if v, ok := htable[myValue]; ok {
			resp = append(resp, arrIdx)
			resp = append(resp, v)
			return resp
		}
	}

	return nil
}
