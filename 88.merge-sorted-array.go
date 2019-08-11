package main

/*
 * @lc app=leetcode id=88 lang=golang
 *
 * [88] Merge Sorted Array
 */
func mergeII(nums1 []int, m int, nums2 []int, n int) {

	for i := m - 1; i >= 0; i-- {
		nums1[i+n] = nums1[i]
	}

	var i, j, k, v int
	for k < m+n {
		if j >= n || i < m && nums1[i+n] <= nums2[j] {
			v = nums1[i+n]
			i++
		} else {
			v = nums2[j]
			j++
		}
		nums1[k] = v
		k++
	}
}

// 与56题重名， 提交时取消下列注释
// func merge(nums1 []int, m int, nums2 []int, n int) {
// 	mergeII(nums1, m, nums2, n)
// }
