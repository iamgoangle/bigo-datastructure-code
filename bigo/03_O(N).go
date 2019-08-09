package main

// Linear

import (
	"fmt"
)

func main() {
	lists := []string{"H", "e", "l", "l", "o"}
	fmt.Println(findChar("H", lists))
}

func findChar(find string, lists []string) bool {
	for _, l := range lists {
		if find == l {
			return true
		}
	}

	return false
}
