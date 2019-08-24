package main

/*
 * @lc app=leetcode id=97 lang=golang
 *
 * [97] Interleaving String
 */
func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s3) != len(s1)+len(s2) {
		return false
	}
	if len(s3) == 0 {
		return true
	}
	if len(s1) == 0 {
		return s2 == s3
	}

	m, n := len(s3)+1, len(s1)+1
	var i, j int
	b := make([][]bool, m)
	for i = 0; i < m; i++ {
		b[i] = make([]bool, n)
	}
	b[0][0] = true
	for i = 1; i < m; i++ {
		b[i][0] = b[i-1][0] && i <= len(s2) && s3[i-1] == s2[i-1]
	}

	for i = 1; i < m; i++ {
		for j = 1; j < n; j++ {
			if b[i-1][j-1] && s3[i-1] == s1[j-1] ||
				b[i-1][j] && 0 <= i-j-1 && i-j-1 < len(s2) && s3[i-1] == s2[i-j-1] {
				b[i][j] = true
			}
		}
	}

	return b[m-1][n-1]
}
