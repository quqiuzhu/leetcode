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

// Matrix 2d slice
type Matrix [][]int

// Dump print matrix
func (m Matrix) Dump() {
	for i := 0; i < len(m); i++ {
		fmt.Print("[")
		for j := 0; j < len(m[i]); j++ {
			fmt.Print(m[i][j], ",")
		}
		fmt.Println("]")
	}
}

// TreeNode binary tree node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
