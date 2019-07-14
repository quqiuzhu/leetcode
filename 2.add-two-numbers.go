package main

/*
 * @lc app=leetcode id=2 lang=golang
 *
 * [2] Add Two Numbers
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbersRemainder(l1 *ListNode, l2 *ListNode, r int) *ListNode {
	var n, n1, n2 *ListNode
	var v int
	if l1 == nil && l2 == nil {
		if r == 0 {
			return nil
		}
		v = r
	} else if l1 == nil {
		v = l2.Val + r
		n2 = l2.Next
	} else if l2 == nil {
		v = l1.Val + r
		n1 = l1.Next
	} else {
		v = l1.Val + l2.Val + r
		n1 = l1.Next
		n2 = l2.Next
	}

	n = &ListNode{Val: v % 10}
	r = v / 10
	n.Next = addTwoNumbersRemainder(n1, n2, r)
	return n
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return addTwoNumbersRemainder(l1, l2, 0)
}
