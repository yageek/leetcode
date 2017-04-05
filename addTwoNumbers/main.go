package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	s := ""
	for n := l; n != nil; n = n.Next {
		s += fmt.Sprintf("%d", n.Val)
	}
	return s
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	node := &ListNode{}
	lhs := l1
	rhs := l2
	carriage := 0
	for {
		if rhs == nil && lhs == nil {
			break
		}
		lhsValue := 0
		rhsValue := 0
		if rhs != nil {
			rhsValue = rhs.Val
		}
		if lhs != nil {
			lhsValue = lhs.Val
		}
		node.Val = (rhsValue + lhsValue + carriage) % 10
		carriage = (rhsValue + lhsValue + carriage) / 10

		node = &ListNode{Next: node}
		rhs = rhs.Next
		lhs = lhs.Next
	}
	return node
}

func main() {
	l1 := &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	fmt.Println(l1)
	l2 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	fmt.Println(l2)
	fmt.Println(addTwoNumbers(l1, l2))
}
