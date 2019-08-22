package main

import (
	"fmt"
	"sort"
)

/*
	Given two strings s and t , write a function to determine if t is an anagram of s.

	Example 1:

	Input: s = "anagram", t = "nagaram"
	Output: true
	Example 2:

	Input: s = "rat", t = "car"
	Output: false
	Note:
	You may assume the string contains only lowercase alphabets.

	Follow up:
	What if the inputs contain unicode characters? How would you adapt your solution to such case?

	https://leetcode.com/explore/featured/card/top-interview-questions-easy/127/strings/882/
*/

func main() {
	s := "anagram"
	t := "nagaram"

	fmt.Println(isAnagram(s, t))
}

func isAnagram(s string, t string) bool {
	stringS := ""
	stringT := ""

	sSort := []int{}
	tSort := []int{}

	if len(s) != len(t) {
		return false
	}

	for i := 0; i < len(s); i++ {
		sSort = append(sSort, int(s[i]))
		tSort = append(tSort, int(t[i]))
	}

	sort.Ints(sSort)
	sort.Ints(tSort)

	for j := 0; j < len(t); j++ {
		stringS += string(sSort[j])
		stringT += string(tSort[j])
	}

	return stringS == stringT
}
