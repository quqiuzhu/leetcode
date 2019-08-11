package main

/*
 * @lc app=leetcode id=86 lang=golang
 *
 * [86] Partition List
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	var it, lh, lt, rh, rt *ListNode
	it = head
	for it != nil {
		if it.Val < x {
			if lh == nil {
				lh = it
				lt = it
			} else {
				lt.Next = it
				lt = it
			}
		} else {
			if rh == nil {
				rh = it
				rt = it
			} else {
				rt.Next = it
				rt = it
			}
		}
		it = it.Next
	}
	if lh == nil {
		lh = rh
	} else {
		lt.Next = rh
	}
	if rt != nil {
		rt.Next = nil
	}
	return lh
}
