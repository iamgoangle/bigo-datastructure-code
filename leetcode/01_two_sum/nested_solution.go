package main

import (
	"fmt"
)

/*
	Time complexity : O(n^2)
	Space complexity : O(1)
*/
func main() {
	nums := []int{2, 7, 11, 15}
	target := 9

	twoNum(nums, target)
}

func twoNum(nums []int, target int) {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			sum := (nums[i] + nums[j])
			fmt.Println(sum)
			if target == sum {
				fmt.Println("Matched index = ", i, j)
				return
			} else {
				fmt.Println("not match")
			}
		}
	}
}
