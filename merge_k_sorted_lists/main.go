package main

import "math"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	root := &ListNode{}

	pointer := root
	for {
		min := math.MaxInt
		minIdx := -1
		for index, list := range lists {
			if list != nil && list.Val < min {
				min = list.Val
				minIdx = index
			}
		}
		if minIdx == -1 {
			break
		}
		lists[minIdx] = lists[minIdx].Next
		pointer.Next = &ListNode{Val: min}
		pointer = pointer.Next

	}
	return root.Next
}
