package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	L1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}

	L2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	// fmt.Printf("%+v", addTwoNumbers(L1, L2))
	print(addTwoNumbers(L1, L2))
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return &ListNode{
		Val: l1.Val + l2.Val,
		Next: &ListNode{
			Val: l1.Next.Val + l2.Next.Val,
			Next: &ListNode{
				Val:  l1.Next.Next.Val + l2.Next.Next.Val,
				Next: nil,
			},
		},
	}
}

func print(l *ListNode) {
	fmt.Println(l.Val)
	nextNode := l.Next

	if nextNode != nil {
		print(nextNode)
	}
}
