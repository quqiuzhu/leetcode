package main

/*
 * @lc app=leetcode id=21 lang=golang
 *
 * [21] Merge Two Sorted Lists
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	const MaxInt = int(^uint(0) >> 1)

	it1, it2 := l1, l2
	var head, tail, it *ListNode
	var v1, v2 int
	for it1 != nil || it2 != nil {
		v1, v2 = MaxInt, MaxInt
		if it1 != nil {
			v1 = it1.Val
		}
		if it2 != nil {
			v2 = it2.Val
		}

		it = nil
		if v1 <= v2 {
			it = it1
			it1 = it1.Next
		} else {
			it = it2
			it2 = it2.Next
		}

		if head == nil {
			head = it
			tail = it
		} else {
			tail.Next = it
			tail = it
		}
	}
	return head
}
