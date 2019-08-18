package main

import (
	"fmt"
	"strings"
)

/*
  Input: "pwwkew"
  Output: 3
  Explanation: The answer is "wke", with the length of 3.
               Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
*/

func main() {
	// input := "pwwkew"
	// input := "abcabcbb"
	input := "bbbbb"

	result := ""
	count := 0

	for i := 0; i < len(input); i++ {
		if !strings.Contains(result, string(input[i])) {
			result = result + string(input[i])
		} else {
			fmt.Println(result)

			if len(result) > count {
				count = len(result)
			}
			result = string(input[i])
		}
	}

	fmt.Println(count)

}
