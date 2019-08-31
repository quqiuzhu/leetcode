package main

/*
 * @lc app=leetcode id=141 lang=golang
 *
 * [141] Linked List Cycle
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	for head != nil {
		if head == head.Next {
			return true
		}
		head, head.Next = head.Next, head
	}
	return false
}
