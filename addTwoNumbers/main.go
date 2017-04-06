package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	s := "List:"
	for n := l; n != nil; n = n.Next {
		s += fmt.Sprintf("%d", n.Val)
	}
	return s
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var previous *ListNode
	var first *ListNode
	lhs := l1
	rhs := l2
	carriage := 0
	for {
		if rhs == nil && lhs == nil && carriage == 0 {
			break
		}
		lhsValue := 0
		rhsValue := 0
		if rhs != nil {
			rhsValue = rhs.Val
			rhs = rhs.Next
		}
		if lhs != nil {
			lhsValue = lhs.Val
			lhs = lhs.Next
		}

		sum := (rhsValue + lhsValue + carriage) % 10
		carriage = (rhsValue + lhsValue + carriage) / 10
		node := &ListNode{Val: sum}
		if previous != nil {
			previous.Next = node
		} else {
			first = node
		}
		previous = node
	}
	return first
}

func main() {
	l1 := &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	l2 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	fmt.Println(addTwoNumbers(l1, l2))

	fmt.Println(addTwoNumbers(&ListNode{5, nil}, &ListNode{5, nil}))
}
