package main

/*
 * @lc app=leetcode id=154 lang=golang
 *
 * [154] Find Minimum in Rotated Sorted Array II
 */
func findMinII(nums []int) int {
	var last, min, i int
	last = nums[len(nums)-1]
	min = last

	for i = 0; i < len(nums)-1; i++ {
		if nums[i] != last {
			break
		}
	}
	nums = nums[i:]

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

// 与153题重名， 提交时取消下列注释
// func findMin(nums []int) int {
// 	return findMinII(nums)
// }
