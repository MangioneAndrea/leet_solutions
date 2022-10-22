package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	var s = ""
	for l != nil {
		s += fmt.Sprintf("%d", l.Val)
		l = l.Next
	}
	return s
}

func add(l1 *ListNode, l2 *ListNode, to *ListNode, port int) {
	var res = 0
	v1 := 0
	v2 := 0

	var n1 *ListNode
	var n2 *ListNode

	if l1 != nil {
		v1 = l1.Val
		n1 = l1.Next
	}
	if l2 != nil {
		v2 = l2.Val
		n2 = l2.Next
	}

	res = v1 + v2 + port

	to.Val = res % 10
	port = res / 10

	if n1 != nil || n2 != nil || port != 0 {
		to.Next = &ListNode{0, nil}
		add(n1, n2, to.Next, port)
	}

}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var Result = &ListNode{0, nil}
	add(l1, l2, Result, 0)

	return Result
}

func main() {
	res := addTwoNumbers(&ListNode{2, &ListNode{4, &ListNode{3, nil}}}, &ListNode{5, &ListNode{6, &ListNode{4, nil}}})
	fmt.Println(res.String())
}
