package main

import (
  "fmt"
)

func main() {
  input := []int{141, 163, 164, 145, 162}
  
  for i:=0; i<len(input); i++ {
    for j:=0;j<len(input)-i-1;j++{
      if input[j] > input[j+1] {
        temp := input[j]
        input[j] = input[j+1]
        input[j+1] = temp
      }
    }
  }
  
  fmt.Println(input)
}