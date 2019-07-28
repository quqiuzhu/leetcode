package main

/*
 * @lc app=leetcode id=73 lang=golang
 *
 * [73] Set Matrix Zeroes
 */
func setZeroes(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	m0 := make([]bool, m)
	n0 := make([]bool, n)

	var i, j int
	for i = 0; i < m; i++ {
		for j = 0; j < n; j++ {
			if matrix[i][j] == 0 {
				m0[i] = true
				n0[j] = true
			}
		}
	}
	for i = 0; i < m; i++ {
		for j = 0; j < n; j++ {
			if m0[i] || n0[j] {
				matrix[i][j] = 0
			}
		}
	}
}
