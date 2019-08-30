package main

/*
 * @lc app=leetcode id=132 lang=golang
 *
 * [132] Palindrome Partitioning II
 */
func isPalindromeIV(s string) bool {
	ok := true
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			ok = false
			break
		}
	}
	return ok
}
func minCut(s string) int {
	if len(s) <= 1 {
		return 0
	}
	l := len(s)
	m := make([]int, l)
	m[0] = 0
	var i, j int
	for i = 1; i < l; i++ {
		if isPalindromeIV(s[:i+1]) {
			m[i] = 0
			continue
		}
		m[i] = m[i-1]
		for j = i - 1; j > 0; j-- {
			if !isPalindromeIV(s[j : i+1]) {
				continue
			}
			if m[j-1] < m[i] {
				m[i] = m[j-1]
			}
		}
		m[i]++
	}
	return m[l-1]
}
