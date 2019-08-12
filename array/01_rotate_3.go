package main

import "fmt"

// expect [5, 6, 7, 1, 2, 3, 4]
func main() {
	lists := []int{1, 2, 3, 4, 5, 6, 7}

	leftRotate(lists, 5, len(lists))
	fmt.Println("Result => ", lists)
}

func leftRotate(lists []int, d, n int) {
	size := n

	// 1, 2, 3, 4, 5
	temp := []int{}
	for i := 0; i < d; i++ {
		temp = append(temp, lists[i])
	}
	fmt.Println("temp", temp)

	// swap
	// 1, 2, 3, 4, 5, 6, 7
	// 6, 2, 3, 4, 5, 6, 7
	// 6, 7, 3, 4, 5, 6, 7
	for k := 0; k < size-d; k++ {
		lists[k] = lists[k+d]
		fmt.Println("lists", lists)
	}

	// rotate temp to tail
	// start from array len - swap len
	swapCount := size - d
	tIdx := 0
	for j := swapCount; j < size; j++ {
		lists[j] = temp[tIdx]
		tIdx++
	}
}
