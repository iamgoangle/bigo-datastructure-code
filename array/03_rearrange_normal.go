package main

import (
	"fmt"
)

func main() {
	lists := []int{-1, -1, 6, 1, 9, 3, 2, -1, 4, -1}
	size := len(lists)
	hashReArrange(lists, size)
}

func hashReArrange(l []int, s int) {
	hash := make(map[int]int)
	for i := 0; i < s; i++ {
		hash[l[i]] = l[i]
	}

	result := []int{}
	for j := 0; j < s; j++ {
		if _, ok := hash[j]; ok {
			result = append(result, j)
		} else {
			result = append(result, -1)
		}
	}

	fmt.Println(result)
}
