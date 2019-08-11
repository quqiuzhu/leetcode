package main

/*
 * @lc app=leetcode id=87 lang=golang
 *
 * [87] Scramble String
 */
func isScramble(s1 string, s2 string) bool {
	if len(s1) != len(s2) || len(s1) == 0 {
		return false
	}

	var m [][][]bool
	var i, j, k, l, n int
	n = len(s1)
	m = make([][][]bool, n)
	for i = 0; i < n; i++ {
		m[i] = make([][]bool, n)
		for j = 0; j < n; j++ {
			m[i][j] = make([]bool, n+1)
			m[i][j][1] = s1[i] == s2[j]
		}
	}

	for l = 2; l <= n; l++ {
		for i = 0; i <= n-l; i++ {
			for j = 0; j <= n-l; j++ {
				for k = 1; k < l; k++ {
					if (m[i][j][k] && m[i+k][j+k][l-k]) ||
						(m[i+k][j][l-k] && m[i][j+l-k][k]) {
						m[i][j][l] = true
					}
				}
			}
		}
	}

	return m[0][0][n]
}
