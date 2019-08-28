package main

import (
	"fmt"
	"strings"
)

func main() {
	// inputs := []string{"leets", "leetcode", "leet", "leeds"}
	inputs := []string{"flower", "flow", "flight"}

	if len(inputs) == 0 {
		fmt.Println("")
	}

	prefix := inputs[0]
	for i := 1; i < len(inputs); i++ {
		// for strings.Index(prefix, inputs[i]) != -1 {
		for strings.HasPrefix(prefix, inputs[i]) {
			prefix = prefix[0 : len(prefix)-1]

			if len(prefix) == 0 {
				return
			}
		}
	}

	fmt.Println(prefix)
}
