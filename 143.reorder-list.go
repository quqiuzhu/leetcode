package main

/*
 * @lc app=leetcode id=143 lang=golang
 *
 * [143] Reorder List
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	if head == nil {
		return
	}

	var it, tail *ListNode
	half, count := 0, 0
	it = head
	for it != nil {
		count++
		it = it.Next
	}
	half = (count + 1) / 2

	count = 0
	it = head
	for it != nil {
		count++
		if count == half {
			tail = it.Next
			it.Next = nil
			break
		}
		it = it.Next
	}

	it = tail
	tail = nil
	for it != nil {
		it.Next, it, tail = tail, it.Next, it
	}

	it = head
	for it != nil {
		if tail == nil {
			break
		}
		it, it.Next, tail, tail.Next = it.Next, tail, tail.Next, it.Next
	}
}
