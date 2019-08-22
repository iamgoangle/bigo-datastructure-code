// To execute Go code, please declare a func main() in a package "main"

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addNodeEnd(newNode, node *ListNode) *ListNode {
	if node == nil {
		return node
	}

	for n := node; n != nil; n = n.Next {
		if n.Next == nil {
			n.Next = newNode
			return node
		}
	}

	return node
}

func deleteNode(val int, node *ListNode) *ListNode {
	if node == nil {
		return node
	}

	n := node
	if n.Val == val {
		n = n.Next
		return n
	}

	for n.Next != nil {
		if n.Next.Val == val {
			n.Next = n.Next.Next // 2 -> 3
			return node
		}

		n = n.Next
	}

	return node
}

func printLinkedList(node *ListNode) {
	for n := node; n != nil; n = n.Next {
		fmt.Println(n)
	}
}

func main() {
	list1 := &ListNode{
		Val:  1,
		Next: nil,
	}

	list2 := &ListNode{
		Val:  2,
		Next: nil,
	}

	list3 := &ListNode{
		Val:  3,
		Next: nil,
	}

	node := list1
	node = addNodeEnd(list2, node)
	node = addNodeEnd(list3, node)
	node = deleteNode(2, node)

	printLinkedList(node)
}
