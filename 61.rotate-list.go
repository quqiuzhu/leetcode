package main

/*
 * @lc app=leetcode id=61 lang=golang
 *
 * [61] Rotate List
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	var it, kPlus1th, tail *ListNode
	it = head
	count := 0
	for it != nil {
		count++
		it = it.Next
	}
	k = k % count

	kPlus1th = head
	it = head
	count = 0
	for it != nil {
		count++
		if count > k+1 {
			kPlus1th = kPlus1th.Next
		}
		tail = it
		it = it.Next
	}

	tail.Next = head
	if kPlus1th.Next != nil {
		it = kPlus1th.Next
	} else {
		it = head
	}

	kPlus1th.Next = nil
	return it
}
