package main

import (
	"fmt"
)

func main() {
	// arrA := []int{7, 1, 5, 2, 3, 6}
	// arrB := []int{3, 8, 6, 20, 7}

	arrA := []int{1, 2, 2, 1}
	arrB := []int{2, 2}

	hMapA := make(map[int]int)
	hMapB := make(map[int]int)
	result := []int{}

	// hash
	for i := 0; i < len(arrA); i++ {
		if _, ok := hMapA[arrA[i]]; !ok {
			hMapA[arrA[i]] = 1
		} else {
			hMapA[arrA[i]] += 1
		}
	}

	for i := 0; i < len(arrB); i++ {
		if _, ok := hMapA[arrB[i]]; ok {
			if hMapA[arrB[i]] > 0 {
				result = append(result, arrB[i])
			}
		}
	}

	// for k, v := range hMapA {
	//   if v > 1 {
	//     result = append(result, k)
	//   }
	// }

	fmt.Printf("%+v\n", hMapA)
	fmt.Printf("%+v\n", hMapB)
	fmt.Printf("%+v\n", result)
}
