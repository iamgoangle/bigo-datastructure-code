// To execute Go code, please declare a func main() in a package "main"

package main

import "fmt"

/*
	Delete Node in a Linked List
	Answer
	Write a function to delete a node (except the tail) in a singly linked list, given only access to that node.

	Given linked list -- head = [4,5,1,9], which looks like following:
	Input: head = [4,5,1,9], node = 5

	https://leetcode.com/explore/featured/card/top-interview-questions-easy/93/linked-list/553/
*/

type NodeList struct {
	Val  int
	Next *NodeList
}

func main() {
	l1 := &NodeList{
		Val:  1,
		Next: nil,
	}

	l2 := &NodeList{
		Val:  5,
		Next: nil,
	}

	l3 := &NodeList{
		Val:  1,
		Next: nil,
	}

	l4 := &NodeList{
		Val:  9,
		Next: nil,
	}

	node := l1
	node = addListEnd(l2, node)
	node = addListEnd(l3, node)
	node = addListEnd(l4, node)

	node = deleteList(5, node)

	printList(node)
}

func printList(node *NodeList) {
	for n := node; n != nil; n = n.Next {
		fmt.Println(n)
	}
}

func addListEnd(newNode, node *NodeList) *NodeList {
	if newNode == nil {
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

func deleteList(val int, node *NodeList) *NodeList {
	if node == nil {
		return nil
	}

	if node.Val == val {
		node.Next = nil
		return node
	}

	// n.Val == 1
	for n := node; n != nil; n = n.Next {
		if n.Next != nil {
			if n.Next.Val == val {
				nextPointer := n.Next.Next
				n.Next = nextPointer
				return node
			}
		}
	}

	return node
}
