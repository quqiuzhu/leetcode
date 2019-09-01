package main

/*
 * @lc app=leetcode id=142 lang=golang
 *
 * [142] Linked List Cycle II
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	fast := head
	slow := head
	for fast != nil && slow != nil {
		slow = slow.Next
		if fast.Next == nil {
			break
		}
		fast = fast.Next.Next
		if slow == fast {
			for head != slow {
				head = head.Next
				slow = slow.Next
			}
			return head
		}
	}
	return nil
}
