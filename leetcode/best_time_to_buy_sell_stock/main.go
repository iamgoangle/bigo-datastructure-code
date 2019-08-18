package main

import (
	"fmt"
)

/*
  Input: [7,1,5,3,6,4]
  Output: 7
  Explanation: Buy on day 2 (price = 1) and sell on day 3 (price = 5), profit = 5-1 = 4.
             Then buy on day 4 (price = 3) and sell on day 5 (price = 6), profit = 6-3 = 3.

  https://leetcode.com/explore/featured/card/top-interview-questions-easy/92/array/564/
*/

func main() {
	prices := []int{7, 1, 5, 3, 4, 6}
	// prices := []int{10, 2, 3, 6}

	res := 0

	for i := 0; i < len(prices)-1; i++ {
		if prices[i] < prices[i+1] {
			res = res + prices[i+1] - prices[i]
			fmt.Println("i => ", prices[i])
			fmt.Println("res => ", res)
		}
	}

	fmt.Println(res)

}
