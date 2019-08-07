package main

/*
 * @lc app=leetcode id=81 lang=golang
 *
 * [81] Search in Rotated Sorted Array II
 */
func searchII(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}

	if nums[0] == target {
		return true
	}

	var i int
	for i = len(nums) - 1; i >= 0; i-- {
		if nums[i] != nums[0] {
			break
		}
	}
	nums = nums[0 : i+1]

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
	return r != -1
}

// 与33题重名， 提交时取消注释
// func search(nums []int, target int) bool {
// 	return searchII(nums, target)
// }
