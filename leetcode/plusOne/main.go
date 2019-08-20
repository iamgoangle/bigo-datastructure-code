package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3}
	nums[len(nums)-1] = nums[len(nums)-1] + 1

	fmt.Println(nums)
}
