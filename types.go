package main

import "fmt"

// ListNode linklist item
type ListNode struct {
	Val  int
	Next *ListNode
}

// Dump print list
func (l *ListNode) Dump() {
	it := l
	for it != nil {
		fmt.Print(it.Val, "->")
		it = it.Next
	}
	fmt.Println()
}

// NewListNode from slice
func NewListNode(items []int) *ListNode {
	var head, it *ListNode
	for i := 0; i < len(items); i++ {
		if head == nil {
			head = &ListNode{Val: items[i]}
			it = head
		} else {
			it.Next = &ListNode{Val: items[i]}
			it = it.Next
		}
	}
	return head
}
