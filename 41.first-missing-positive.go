package main

/*
 * @lc app=leetcode id=41 lang=golang
 *
 * [41] First Missing Positive
 */
func firstMissingPositive(nums []int) int {
	begin, end, l, i := 0, len(nums)-1, len(nums), 0
	for begin <= end {
		i = nums[begin] - 1
		if nums[begin] == begin+1 {
			begin++
		} else if nums[begin] <= 0 || nums[begin] > l {
			nums[begin], nums[end] = nums[end], nums[begin]
			end--
		} else if nums[begin] == nums[i] {
			begin++
		} else {
			nums[begin], nums[i] = nums[i], nums[begin]
		}
	}
	for i = 0; i < l; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return l + 1
}
