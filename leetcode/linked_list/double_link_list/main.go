package main

import (
	"fmt"
)

type NodeList struct {
	data int
	next *NodeList
	prev *NodeList
}

type List struct {
	head *NodeList
	tail *NodeList
}

func main() {
	list := &List{}
	insert(1, list)
	insert(2, list)
	insert(3, list)
	insert(4, list)

	fmt.Printf("head = %+v\n", list.head)

	fmt.Printf("%+v\n", list.head)
	fmt.Printf("%+v\n", list.head.next)
	fmt.Printf("%+v\n", list.head.next.next)

	fmt.Printf("tail = %+v\n", list.tail)

	print(list)
}

func insert(v int, list *List) *List {
	newNode := &NodeList{
		data: v,
	}

	if list.head == nil {
		list.head = newNode
		return list
	}

	for l := list.head; l != nil; l = l.next {
		if l.next == nil {
			l.next = newNode
			l.next.prev = l // next prev is mine

			list.tail = newNode // tail is last node
			return list
		}
	}

	return list
}

func print(list *List) {
	fmt.Println("Summary")
	for l := list.head; l != nil; l = l.next {
		if l.next != nil {
			fmt.Printf("%v -> %v\n", l.data, l.next.data)
		}
	}
}
