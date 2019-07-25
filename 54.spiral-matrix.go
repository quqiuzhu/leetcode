package main

/*
 * @lc app=leetcode id=54 lang=golang
 *
 * [54] Spiral Matrix
 */

func index(c, index, m, n int) (int, int) {
	mw, nw := m-2*c-1, n-2*c-1
	var i, j int
	if index < nw {
		i = c
		j = c
		if nw != 0 {
			j += index % nw
		}
	} else if index < mw+nw {
		j = n - 1 - c
		i = c
		if mw != 0 {
			i += (index - nw) % mw
		}
	} else if index < 2*nw+mw {
		i = m - 1 - c
		j = c + nw
		if nw != 0 {
			j -= (index - nw - mw) % nw
		}
	} else if index < 2*nw+2*mw {
		j = c
		i = c + mw
		if mw != 0 {
			i -= (index - (2*nw + mw)) % mw
		}
	} else {
		i = c
		j = c
	}
	return i, j
}

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}

	var rs []int
	var c, cl, k, i, j int
	m, n := len(matrix), len(matrix[0])
	for c = 0; c < (m+1)/2; c++ {
		cl = 2 * (n - 1 - 2*c + m - 1 - 2*c)
		for k = 0; k < cl && len(rs) < m*n; k++ {
			i, j = index(c, k, m, n)
			// fmt.Println(i, j)
			rs = append(rs, matrix[i][j])
		}
	}
	if len(rs) < m*n {
		rs = append(rs, matrix[c-1][c-1])
	}
	return rs
}
