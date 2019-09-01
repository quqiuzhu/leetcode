package main

import (
	"sort"
)

/*
 * @lc app=leetcode id=148 lang=golang
 *
 * [148] Sort List
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	var it *ListNode
	var values []int
	it = head
	for it != nil {
		values = append(values, it.Val)
		it = it.Next
	}

	sort.Sort(sort.IntSlice(values))

	i := 0
	it = head
	for it != nil {
		it.Val = values[i]
		i++
		it = it.Next
	}
	return head
}
