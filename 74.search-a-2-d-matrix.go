package main

/*
 * @lc app=leetcode id=74 lang=golang
 *
 * [74] Search a 2D Matrix
 */
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}

	m, n := len(matrix), len(matrix[0])
	var begin, mid, end, i, j int
	end = m*n - 1
	for begin <= end {
		mid = (begin + end) / 2
		i = mid / n
		j = mid % n
		if matrix[i][j] == target {
			return true
		}

		if matrix[i][j] > target {
			end = mid - 1
		} else {
			begin = mid + 1
		}
	}
	return false
}
