package main

/*
 * @lc app=leetcode id=96 lang=golang
 *
 * [96] Unique Binary Search Trees
 */
func numTrees(n int) int {
	if n == 0 {
		return 0
	}
	nums := make([]int, n+1)
	nums[0] = 1
	nums[1] = 1
	var i, j int
	for i = 2; i <= n; i++ {
		for j = 0; j < i; j++ {
			nums[i] += nums[j] * nums[i-1-j]
		}
	}
	return nums[n]
}
