package main

/*
 * @lc app=leetcode id=83 lang=golang
 *
 * [83] Remove Duplicates from Sorted List
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	var it, p *ListNode
	it, p = head, head
	for it != nil {
		if it.Val != p.Val {
			p.Next = it
			p = it
		}
		it = it.Next
	}
	if p != nil {
		p.Next = nil
	}
	return head
}
