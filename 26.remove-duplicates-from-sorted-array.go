package main

/*
 * @lc app=leetcode id=26 lang=golang
 *
 * [26] Remove Duplicates from Sorted Array
 */
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	x, i, r := nums[0]-1, 0, 0
	for i = 0; i < len(nums); i++ {
		if nums[i] != x {
			x = nums[i]
			nums[r] = x
			r++
		}
	}
	return r
}
