package main

import (
	"fmt"
)

/*
	Given a string, find the first non-repeating character in it and return it's index. If it doesn't exist, return -1.

	Examples:

	s = "leetcode"
	return 0.

	s = "loveleetcode",
	return 2.
*/

func main() {
	// s := "leetcode"
	s := "loveleetcode"

	fmt.Println(firstUniqChar(s))
}

func firstUniqChar(s string) int {
	hTable := make(map[string]int)
	for i := 0; i < len(s); i++ {
		hTable[string(s[i])] += 1
	}

	fmt.Println(hTable)

	for j := 0; j < len(s); j++ {
		if _, ok := hTable[string(s[j])]; ok {
			if hTable[string(s[j])] == 1 {
				return j
			}
		}
	}

	return -1
}
