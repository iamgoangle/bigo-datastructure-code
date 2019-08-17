package main

import (
	"container/list"
	"fmt"
)

func main() {
	aList := list.New()
	aList.PushBack(2)
	aList.PushBack(4)
	aList.PushBack(3)

	bList := list.New()
	bList.PushBack(5)
	bList.PushBack(6)
	bList.PushBack(4)

	resultList := list.New()

	a := aList.Front()
	b := bList.Front()
	carry := 0

	for a != nil && b != nil {
		fmt.Println("a", a.Value.(int))
		fmt.Println("b", b.Value.(int))

		cal := a.Value.(int) + b.Value.(int)
		cal = cal + carry
		mod := cal % 10
		divide := cal / 10

		carry = divide
		fmt.Println("mod", mod)
		fmt.Println("divide", divide)

		if mod == 0 {
			resultList.PushBack(mod)
		} else {
			resultList.PushBack(cal)
		}

		a = a.Next()
		b = b.Next()
	}

	for r := resultList.Front(); r != nil; r = r.Next() {
		fmt.Printf("%v -> ", r.Value.(int))
	}
}
