package main

/*
 * @lc app=leetcode id=72 lang=golang
 *
 * [72] Edit Distance
 */
func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	if m == 0 {
		return n
	}
	if n == 0 {
		return m
	}

	d := make([][]int, m)
	for i := 0; i < m; i++ {
		d[i] = make([]int, n)
	}

	d[0][0] = 1
	if word1[0] == word2[0] {
		d[0][0] = 0
	}

	for i := 1; i < m; i++ {
		if word1[i] == word2[0] {
			d[i][0] = i
		} else {
			d[i][0] = d[i-1][0] + 1
		}
	}

	for i := 1; i < n; i++ {
		if word1[0] == word2[i] {
			d[0][i] = i
		} else {
			d[0][i] = d[0][i-1] + 1
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			d[i][j] = d[i-1][j-1]
			if word1[i] == word2[j] {
				continue
			}

			if d[i-1][j] < d[i][j] {
				d[i][j] = d[i-1][j]
			}
			if d[i][j-1] < d[i][j] {
				d[i][j] = d[i][j-1]
			}
			d[i][j]++
		}
	}

	return d[m-1][n-1]
}
