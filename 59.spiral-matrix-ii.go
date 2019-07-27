package main

/*
 * @lc app=leetcode id=59 lang=golang
 *
 * [59] Spiral Matrix II
 */

func idxII(c, index, l int) (int, int) {
	var i, j, k int
	width := l - 2*c - 1
	k = index / width
	if k == 0 {
		i = c
		j = c + index%width
	} else if k == 1 {
		j = l - 1 - c
		i = c + index%width
	} else if k == 2 {
		i = l - 1 - c
		j = c + width - (index % width)
	} else {
		j = c
		i = c + width - (index % width)
	}
	return i, j
}

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}

	var c, index, width, cl, i, j, counter int
	counter++
	for c = 0; c < n/2; c++ {
		width = n - 2*c - 1
		cl = width * 4
		for index = 0; index < cl; index++ {
			i, j = idxII(c, index, n)
			matrix[i][j] = counter
			counter++
		}
	}
	if counter == n*n {
		matrix[n/2][n/2] = counter
	}
	return matrix
}
