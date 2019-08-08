package main

/*
 * @lc app=leetcode id=82 lang=golang
 *
 * [82] Remove Duplicates from Sorted List II
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	var candidate, it, h, t *ListNode
	var count int

	it = head
	for {
		if candidate == nil {
			candidate = it
		}
		if it != nil && candidate.Val == it.Val {
			count++
			it = it.Next
			continue
		}
		if count == 1 {
			candidate.Next = nil
			if h == nil {
				h = candidate
				t = candidate
			} else {
				t.Next = candidate
				t = t.Next
			}
		}
		if it == nil {
			break
		}
		candidate = it
		count = 1
		it = it.Next
	}

	return h
}
