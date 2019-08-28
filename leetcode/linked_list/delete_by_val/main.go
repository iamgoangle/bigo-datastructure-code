package main

import (
	"fmt"
)

/*
	Given a linked list, remove the n-th node from the end of list and return its head.

	Example:

	Given linked list: 1->2->3->4->5, and n = 2.

	After removing the second node from the end, the linked list becomes 1->2->3->5.
	Note:

	Given n will always be valid.

	Follow up:

	Could you do this in one pass?
	https://leetcode.com/explore/featured/card/top-interview-questions-easy/93/linked-list/603/
*/
type NodeList struct {
	data int
	next *NodeList
}

func main() {
	list := &NodeList{}
	insert(1, list)
	insert(2, list)
	insert(3, list)
	insert(4, list)
	insert(5, list)

	removeNthFromEnd(list, 4)

	display(list)
}

func insert(data int, list *NodeList) *NodeList {
	newNode := &NodeList{
		data: data,
		next: nil,
	}

	for l := list; l != nil; l = l.next {
		if l.next == nil {
			l.next = newNode
			return list
		}
	}

	return list
}

func removeNthFromEnd(list *NodeList, find int) *NodeList {
	for l := list; l != nil; l = l.next {
		if l.next.data == find {
			l.next = l.next.next
			return list
		}
	}

	return list
}

func display(list *NodeList) {
	for l := list; l != nil; l = l.next {
		fmt.Printf("%+v\n", l)
	}
}
