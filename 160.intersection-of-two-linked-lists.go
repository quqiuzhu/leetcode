package main

/*
 * @lc app=leetcode id=160 lang=golang
 *
 * [160] Intersection of Two Linked Lists
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	var onceA, onceB bool
	itA := headA
	itB := headB
	for {
		if itA == nil && !onceA {
			onceA = true
			itA = headB
		}
		if itB == nil && !onceB {
			onceB = true
			itB = headA
		}
		if itA == nil || itB == nil {
			break
		}
		if itA == itB {
			return itA
		}
		itA = itA.Next
		itB = itB.Next
	}

	return nil
}
