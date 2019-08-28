package main

import (
	"fmt"
	"strconv"
)

/*
	https://www.facebook.com/careers/life/sample_interview_questions/?redirect_id=Af_JJya3y6WseT16cfGEN5o3PGdD61i3lRzYJqpgbqb31xN9AFWJ1to7EDzeXhBKFuQ&loop_id=Af_YZp68Vq3o1RQ--T0QABdkCoW58cPBz76zxgeMefRQ9_9CSOH8JXiZfBUW1iLNFp0
  Implement a function that outputs the Look and Say sequence:

  1
  11
  21
  1211
  111221
  312211
  13112221
  1113213211
  31131211131221
  13211311123113112211
*/

func main() {
	printNumberAndSay(3)
}

func printNumberAndSay(n int) {
	//   if n == 1 {
	//     fmt.Println("1")
	//     return
	//   }

	//   if n == 2 {
	//     fmt.Println("11")
	//     return
	//   }

	str := "1"

	for i := 1; i <= n; i++ {
		str += "-"
		count := 1
		tmp := ""

		for j := 1; j < len(str); j++ {
			if string(str[j]) != string(str[j-1]) {
				tmp += strconv.Itoa(count) + string(str[j-1])
				count = 1
			} else {
				count++
			}
		}

		str = tmp
		fmt.Println(str)
	}
}
