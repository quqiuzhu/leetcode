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

// Dump print tree in levels
func (t *TreeNode) Dump() {
	var cur, next []*TreeNode
	cur = append(cur, t)
	for len(cur) > 0 {
		for i := 0; i < len(cur); i++ {
			if cur[i] != nil {
				fmt.Print(cur[i].Val, " ")
				next = append(next, cur[i].Left)
				next = append(next, cur[i].Right)
			} else {
				fmt.Print(-1, " ")
			}
		}
		fmt.Println()
		cur = next
		next = nil
	}
}

// Copy tree copy
func (t *TreeNode) Copy() *TreeNode {
	if t == nil {
		return nil
	}

	copied := &TreeNode{
		Val: t.Val,
	}

	if t.Left != nil {
		copied.Left = t.Left.Copy()
	}

	if t.Right != nil {
		copied.Right = t.Right.Copy()
	}

	return copied
}

func (t *TreeNode) marshalChildren() []int {
	if t == nil {
		return nil
	}
	var result []int
	if t.Left == nil {
		result = append(result, -1)
	} else {
		result = append(result, t.Left.Val)
	}
	if t.Right == nil {
		result = append(result, -1)
	} else {
		result = append(result, t.Right.Val)
	}
	if t.Left != nil {
		rs := t.Left.marshalChildren()
		result = append(result, rs...)
	}
	if t.Right != nil {
		rs := t.Right.marshalChildren()
		result = append(result, rs...)
	}
	return result
}

// Marshal print tree to slice
func (t *TreeNode) Marshal() []int {
	var result []int
	if t == nil {
		result = append(result, -1)
		return result
	}
	result = append(result, t.Val)
	rs := t.marshalChildren()
	result = append(result, rs...)
	return result
}

// NewTreeNode restore tree from slice
func NewTreeNode(tree []int) *TreeNode {
	if len(tree) == 0 || tree[0] < 0 {
		return nil
	}
	var nodes []*TreeNode
	var i, j int
	for i = 0; i < len(tree); i++ {
		if tree[i] < 0 {
			nodes = append(nodes, nil)
		} else {
			nodes = append(nodes, &TreeNode{Val: tree[i]})
		}
	}

	j = 1
	for i = 0; i < len(nodes); i++ {
		if nodes[i] != nil {
			if j < len(nodes) {
				nodes[i].Left = nodes[j]
				j++
			}
			if j < len(nodes) {
				nodes[i].Right = nodes[j]
				j++
			}
		}
	}
	return nodes[0]
}
