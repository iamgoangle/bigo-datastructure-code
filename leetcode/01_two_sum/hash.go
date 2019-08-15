package main

import (
	"fmt"
)

func main() {
	nums := []int{11, 15, 2, 7}
	target := 9
	twoSum(nums, target)
}

func twoSum(nums []int, target int) {
	hash := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		hash[nums[i]] = i
	}

	fmt.Println(hash)

	for j := 0; j < len(hash); j++ {
		cal := target - nums[j] // 9-2 = 7
		fmt.Println(cal)

		if cal > 0 && hash[cal] != j {
			fmt.Println("Index = ", j, hash[cal])
			return
		}
	}
}
