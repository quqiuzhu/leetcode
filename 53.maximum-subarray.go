package main

/*
 * @lc app=leetcode id=53 lang=golang
 *
 * [53] Maximum Subarray
 */
func onePassMaxSubArray(nums []int) int {
	var max, sum int
	maxSet := false
	for i := 0; i < len(nums); i++ {
		if sum < 0 {
			sum = 0
		}
		sum += nums[i]
		if sum > max || !maxSet {
			max = sum
			maxSet = true
		}
	}
	return max
}

func divideConquerMaxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	mid := len(nums) / 2
	m1 := divideConquerMaxSubArray(nums[:mid])
	m2 := divideConquerMaxSubArray(nums[mid:])

	var max, sum int
	maxSet := false
	for i := mid - 1; i >= 0; i-- {
		sum += nums[i]
		if sum > max || !maxSet {
			max = sum
			maxSet = true
		}
	}
	if max < 0 {
		sum = 0
	} else {
		sum = max
	}
	for i := mid; i < len(nums); i++ {
		sum += nums[i]
		if sum > max || !maxSet {
			max = sum
			maxSet = true
		}
	}
	if max < m1 {
		max = m1
	}
	if max < m2 {
		max = m2
	}
	return max
}

func maxSubArray(nums []int) int {
	return divideConquerMaxSubArray(nums)
}
