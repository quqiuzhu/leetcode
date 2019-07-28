package main

/*
 * @lc app=leetcode id=75 lang=golang
 *
 * [75] Sort Colors
 */
func sortColors(nums []int) {
	var red, white, i int
	for i = 0; i < len(nums); i++ {
		if nums[i] == 0 {
			red++
		} else if nums[i] == 1 {
			white++
		}
	}

	white += red

	for i = 0; i < len(nums); i++ {
		if i < red {
			nums[i] = 0
		} else if i < white {
			nums[i] = 1
		} else {
			nums[i] = 2
		}
	}
}
