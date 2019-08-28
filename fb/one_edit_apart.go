// To execute Go code, please declare a func main() in a package "main"

package main

import (
	"fmt"
	// "strings"
)

func main() {
	fmt.Println(OneEditApart("cat", "cut"))
	// fmt.Println(OneEditApart("cat", "cats"))
	// fmt.Println(OneEditApart("catdf", "catss"))
	// fmt.Println(OneEditApart("cat", "at"))
}

func OneEditApart(s1, s2 string) bool {
	// convert to array
	str1 := []string{}
	str2 := []string{}

	for i := 0; i < len(s1); i++ {
		str1 = append(str1, string(s1[i]))
	}

	for i := 0; i < len(s2); i++ {
		str2 = append(str2, string(s2[i]))
	}

	if len(str2) > len(str1) {
		// swap
		tmp := str1
		str1 = str2
		str2 = tmp
	}

	operation := 0
	if len(str1) == len(str2) {
		for i := 0; i < len(str1); i++ {
			if str1[i] != str2[i] {
				operation++
			}

			if operation > 1 {
				return false
			}
		}
	} else {
		fmt.Println(str1)
		fmt.Println(str2)
		for i := 0; i < len(str1); i++ {

			small := func(i, l int) string {
				if i > l-1 {
					return "$"
				} else {
					return str2[i]
				}
			}

			small2 := small(i, len(str2))

			fmt.Println("str1", str1[i])
			fmt.Println("small", small2)

			if str1[i] != small2 {
				operation++
				if operation > 1 {
					return false
				}
			}
		}
	}

	return true
}
