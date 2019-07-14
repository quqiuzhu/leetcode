package main

/*
 * @lc app=leetcode id=45 lang=golang
 *
 * [45] Jump Game II
 */
func jump(nums []int) int {
	i, j, l := 0, 0, len(nums)
	steps := make([]int, l)
	for i = 0; i < l; i++ {
		for j = i + 1; j < l && j <= i+nums[i]; j++ {
			if steps[j] == 0 || steps[i]+1 < steps[j] {
				steps[j] = steps[i] + 1
			}
		}
	}
	return steps[l-1]
}
