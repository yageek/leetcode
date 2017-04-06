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
	head := &ListNode{}
	current := head
	lhs := l1
	rhs := l2
	carriage := 0
	for rhs != nil && lhs != nil {
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
		sum := rhsValue + lhsValue + carriage
		carriage = sum / 10
		current.Next = &ListNode{Val: sum % 10}
		current = current.Next
	}
	if carriage > 0 {
		current.Next = &ListNode{Val: carriage}
	}
	return head.Next
}

func main() {
	l1 := &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	l2 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	fmt.Println(addTwoNumbers(l1, l2))

	fmt.Println(addTwoNumbers(&ListNode{5, nil}, &ListNode{5, nil}))
}
