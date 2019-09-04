package main

/*
 * @lc app=leetcode id=162 lang=golang
 *
 * [162] Find Peak Element
 */
func findPeakElement(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	peak := 0
	begin, mid, end := 0, 0, len(nums)-1
	if nums[begin] > nums[begin+1] {
		return begin
	}
	if nums[end] > nums[end-1] {
		return end
	}

	for begin <= end {
		mid = (begin + end) / 2
		if mid == 0 || mid == len(nums)-1 {
			break
		}
		if nums[mid-1] < nums[mid] && nums[mid] > nums[mid+1] {
			peak = mid
			break
		}
		if nums[mid-1] < nums[mid] {
			begin = mid + 1
			peak = begin
		} else {
			end = mid - 1
			peak = end
		}
	}
	return peak
}
