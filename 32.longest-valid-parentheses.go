package main

/*
 * @lc app=leetcode id=32 lang=golang
 *
 * [32] Longest Valid Parentheses
 */
func longestValidParentheses(s string) int {
	var stack []int

	max, i, last := 0, 0, -1
	for i = 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else if len(stack) == 0 {
			last = i
		} else if len(stack) == 1 {
			stack = nil
			if i-last > max {
				max = i - last
			}
		} else {
			stack = stack[0 : len(stack)-1]
			if i-stack[len(stack)-1] > max {
				max = i - stack[len(stack)-1]
			}
		}
	}

	return max
}
