package main

/*
 * @lc app=leetcode id=92 lang=golang
 *
 * [92] Reverse Linked List II
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	var h, t, it, tail, next *ListNode
	count := 1
	tail = head
	it = head
	for count <= n {
		next = it.Next
		if count >= m {
			if h == nil {
				h = it
				t = it
			} else {
				it.Next = h
				h = it
			}
		} else {
			tail = it
		}
		count++
		it = next
	}

	if tail != t {
		tail.Next = h
	} else {
		head = h
	}
	t.Next = next
	return head
}
