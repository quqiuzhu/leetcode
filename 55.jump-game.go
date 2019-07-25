package main

/*
 * @lc app=leetcode id=55 lang=golang
 *
 * [55] Jump Game
 */
func canJump(nums []int) bool {
	reach, maxIndex := 0, len(nums)-1
	for i := 0; i < len(nums); i++ {
		if i > reach {
			return false
		}
		if i+nums[i] > reach {
			reach = i + nums[i]
		}
		if reach >= maxIndex {
			break
		}
	}
	return reach >= maxIndex
}
