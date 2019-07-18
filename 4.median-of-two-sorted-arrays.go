package main

/*
 * @lc app=leetcode id=4 lang=golang
 *
 * [4] Median of Two Sorted Arrays
 */

func findKth(nums1 []int, nums2 []int, k int) int {
	if len(nums1) < len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	l1, l2 := len(nums1), len(nums2)
	if l1 > k {
		return findKth(nums1[:k], nums2, k)
	}

	if l2 == 0 {
		return nums1[k-1]
	}

	if k == 1 {
		if nums1[0] > nums2[0] {
			return nums2[0]
		}
		return nums1[0]
	}

	m2 := k / 2
	if m2 > l2 {
		m2 = l2
	}
	m1 := k - m2

	if nums1[m1-1] == nums2[m2-1] {
		return nums1[m1-1]
	}

	if nums1[m1-1] < nums2[m2-1] {
		return findKth(nums1[m1:], nums2, k-m1)
	}

	return findKth(nums1, nums2[m2:], k-m2)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	l := len(nums1) + len(nums2)
	m1 := findKth(nums1, nums2, (l+1)/2)
	if l%2 == 1 {
		return float64(m1)
	}
	m2 := findKth(nums1, nums2, (l+1)/2+1)
	m := float64(m1 + m2)
	return m / 2
}
