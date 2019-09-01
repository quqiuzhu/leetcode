package main

/*
 * @lc app=leetcode id=147 lang=golang
 *
 * [147] Insertion Sort List
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertionSortListWithNode(head, node *ListNode) *ListNode {
	if head == nil {
		return node
	}
	if head.Val > node.Val {
		node.Next = head
		return node
	}
	var it, pre *ListNode
	pre = head
	it = head.Next
	for it != nil {
		if it.Val >= node.Val {
			break
		}
		pre = pre.Next
		it = it.Next
	}
	pre.Next = node
	node.Next = it
	return head
}

func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var it, next *ListNode
	it = head
	head = nil
	for it != nil {
		next = it.Next
		it.Next = nil
		head = insertionSortListWithNode(head, it)
		it = next
	}
	return head
}
