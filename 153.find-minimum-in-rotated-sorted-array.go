package main

/*
 * @lc app=leetcode id=153 lang=golang
 *
 * [153] Find Minimum in Rotated Sorted Array
 */
func findMin(nums []int) int {
	last := nums[len(nums)-1]
	min := last
	begin, mid, end := 0, 0, len(nums)-1
	for begin <= end {
		mid = (begin + end) / 2
		if nums[mid] < min {
			min = nums[mid]
		}

		if nums[mid] > last {
			begin = mid + 1
		} else {
			end = mid - 1
		}
	}
	return min
}
