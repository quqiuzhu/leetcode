package main

/*
 * @lc app=leetcode id=19 lang=golang
 *
 * [19] Remove Nth Node From End of List
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	it := head
	i := 0
	var nth *ListNode
	for it != nil {
		if i > n {
			nth = nth.Next
		} else if i == n {
			nth = head
		}
		i++
		it = it.Next
	}
	if i == n {
		return head.Next
	}

	if nth.Next != nil {
		t := nth.Next.Next
		nth.Next.Next = nil
		nth.Next = t
	}
	return head
}
