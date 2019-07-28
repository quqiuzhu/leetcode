package main

/*
 * @lc app=leetcode id=70 lang=golang
 *
 * [70] Climbing Stairs
 */
func climbStairs(n int) int {
	if n <= 1 {
		return n
	}

	var steps, steps1, steps2 int
	steps2 = 1
	steps1 = 2
	for i := 2; i < n; i++ {
		steps = steps1 + steps2
		steps2 = steps1
		steps1 = steps
	}
	return steps
}
