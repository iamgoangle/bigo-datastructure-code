package main

import (
	"fmt"
)

/*
  Remove duplicate array
  https://leetcode.com/explore/featured/card/top-interview-questions-easy/92/array/727/
*/

func main() {
	nums := []int{1, 1, 2}
	hTable := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		hTable[nums[i]] = hTable[i] + 1
	}

	// fmt.Printf("%v", hTable)

	result := []int{}
	for k, _ := range hTable {
		result = append(result, k)
	}

	fmt.Println(result)
}
