package main

/*
 * @lc app=leetcode id=109 lang=golang
 *
 * [109] Convert Sorted List to Binary Search Tree
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func bottomUpBST(l *ListNode, begin, end int) (*ListNode, *TreeNode) {
	if begin > end {
		return l, nil
	}
	m := (begin + end) / 2
	l, left := bottomUpBST(l, begin, m-1)
	root := &TreeNode{
		Val:  l.Val,
		Left: left,
	}
	l = l.Next
	l, right := bottomUpBST(l, m+1, end)
	root.Right = right
	return l, root
}

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	var length int
	it := head
	for it != nil {
		length++
		it = it.Next
	}
	_, root := bottomUpBST(head, 0, length-1)
	return root
}
