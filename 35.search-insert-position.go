package main

/*
 * @lc app=leetcode id=35 lang=golang
 *
 * [35] Search Insert Position
 */
func searchInsert(nums []int, target int) int {
	r := 0
	begin, mid, end := 0, 0, len(nums)-1
	for begin <= end {
		mid = (begin + end) / 2
		if nums[mid] == target {
			r = mid
			break
		}
		if nums[mid] > target {
			end = mid - 1
		} else {
			begin = mid + 1
		}
		r = begin
	}
	return r
}
