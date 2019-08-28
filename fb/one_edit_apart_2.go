package main

import (
	"fmt"
)

func main() {
	fmt.Println(oneEditApart("cat", "dog"))
	fmt.Println(oneEditApart("cat", "cats"))
	fmt.Println(oneEditApart("cat", "cut"))
	fmt.Println(oneEditApart("cat", "custaa"))
	fmt.Println(oneEditApart("cat", "at"))
	fmt.Println(oneEditApart("cat", "act"))
}

func oneEditApart(a, b string) bool {
	fmt.Println(a, "<=>", b)

	arrA := []string{}
	arrB := []string{}

	if len(a) == 0 || len(b) == 0 {
		return false
	}

	if len(a)-len(b) > 1 || len(a)-len(b) < -1 {
		return false
	}

	lenN := 0
	if len(a) > len(b) {
		lenN = len(a)
	} else {
		lenN = len(b)
	}

	// create array
	diff := 0
	for i := 0; i < lenN; i++ {
		if i > len(a)-1 {
			arrA = append(arrA, "$")
			arrB = append(arrB, string(b[i]))
		} else if i > len(b)-1 {
			arrA = append(arrA, string(a[i]))
			arrB = append(arrB, "$")
		} else {
			arrA = append(arrA, string(a[i]))
			arrB = append(arrB, string(b[i]))
		}
	}

	// compare
	for j := 0; j < lenN; j++ {
		if arrA[j] != arrB[j] {
			diff += 1
		}
	}

	return diff <= 1
}
