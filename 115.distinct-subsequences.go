package main

/*
 * @lc app=leetcode id=115 lang=golang
 *
 * [115] Distinct Subsequences
 */
func numDistinct(s string, t string) int {
	m, n := len(t), len(s)
	if m == 0 || n == 0 {
		return 0
	}
	var i, j, count int
	dp := make([][]int, m)
	for i = 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	for j = 0; j < n; j++ {
		if s[j] == t[0] {
			count++
		}
		dp[0][j] = count
	}
	for i = 1; i < m; i++ {
		count = 0
		for j = 1; j < n; j++ {
			if t[i] == s[j] {
				count += dp[i-1][j-1]
			}
			dp[i][j] = count
		}
	}
	return dp[m-1][n-1]
}
