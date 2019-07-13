package main

/*
 * @lc app=leetcode id=34 lang=golang
 *
 * [34] Find First and Last Position of Element in Sorted Array
 */
func searchRange(nums []int, target int) []int {
	min, max := -1, -1
	begin, mid, end := 0, 0, len(nums)-1
	for begin <= end {
		mid = (begin + end) / 2
		if nums[mid] == target {
			min = mid
		}
		if nums[mid] >= target {
			end = mid - 1
		} else {
			begin = mid + 1
		}
	}

	if min == -1 {
		return []int{-1, -1}
	}

	begin, end = 0, len(nums)-1
	for begin <= end {
		mid = (begin + end) / 2
		if nums[mid] == target {
			max = mid
		}
		if nums[mid] > target {
			end = mid - 1
		} else {
			begin = mid + 1
		}
	}

	return []int{min, max}
}
