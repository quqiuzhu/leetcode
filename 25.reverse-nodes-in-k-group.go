package main

/*
 * @lc app=leetcode id=25 lang=golang
 *
 * [25] Reverse Nodes in k-Group
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || k <= 1 {
		return head
	}

	it, i := head, 0
	for ; it != nil && i < k; i++ {
		it = it.Next
	}
	if i < k {
		return head
	}

	nextGroup := it
	var next *ListNode
	p, it := head, head.Next
	for i = 0; i < k-1; i++ {
		next = it.Next
		it.Next = p
		p = it
		it = next
	}

	head.Next = reverseKGroup(nextGroup, k)
	return p
}
