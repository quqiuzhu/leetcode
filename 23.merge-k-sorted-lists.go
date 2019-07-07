package main

/*
 * @lc app=leetcode id=23 lang=golang
 *
 * [23] Merge k Sorted Lists
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func allNil(lists []*ListNode) bool {
	r := true
	for _, l := range lists {
		if l != nil {
			r = false
			break
		}
	}
	return r
}

func forceMergeKLists(lists []*ListNode) *ListNode {
	const MaxInt = int(^uint(0) >> 1)

	var head, tail, it *ListNode
	ll := len(lists)
	its := make([]*ListNode, ll)
	copy(its, lists)
	v, i, j := MaxInt, 0, 0
	for !allNil(its) {
		for i = 0; i < ll; i++ {
			if its[i] != nil && its[i].Val < v {
				v = its[i].Val
				j = i
			}
		}

		it = its[j]
		its[j] = its[j].Next

		if head == nil {
			head = it
			tail = it
		} else {
			tail.Next = it
			tail = it
		}

		v = MaxInt
	}
	return head
}

func recursiveMergeTwoLists(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = recursiveMergeTwoLists(l1.Next, l2)
		return l1
	}
	l2.Next = recursiveMergeTwoLists(l1, l2.Next)
	return l2
}

func binaryMergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	i, mid, end := 0, 0, len(lists)-1

	for end > 0 {
		mid = (end - 1) / 2
		for i = 0; i <= mid; i++ {
			lists[i] = recursiveMergeTwoLists(lists[i], lists[end-i])
		}
		end = end / 2
	}
	return lists[0]
}

func mergeKLists(lists []*ListNode) *ListNode {
	return binaryMergeKLists(lists)
}
