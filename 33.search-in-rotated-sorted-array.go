package main

/*
 * @lc app=leetcode id=33 lang=golang
 *
 * [33] Search in Rotated Sorted Array
 */
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	r := -1
	begin, mid, end := 0, 0, len(nums)-1
	for begin <= end {
		mid = (begin + end) / 2
		if nums[mid] == target {
			r = mid
			break
		}
		if nums[mid] < nums[0] && (target < nums[mid] || target >= nums[0]) ||
			nums[mid] >= nums[0] && (target < nums[mid] && target >= nums[0]) {
			end = mid - 1
		} else {
			begin = mid + 1
		}
	}
	return r
}
