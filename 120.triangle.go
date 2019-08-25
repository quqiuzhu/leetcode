package main

/*
 * @lc app=leetcode id=120 lang=golang
 *
 * [120] Triangle
 */
func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}

	var i, j, total int
	for i = 1; i < len(triangle); i++ {
		for j = 0; j < len(triangle[i]); j++ {
			if j == len(triangle[i])-1 || j != 0 && triangle[i-1][j-1] < triangle[i-1][j] {
				triangle[i][j] += triangle[i-1][j-1]
			} else {
				triangle[i][j] += triangle[i-1][j]
			}
		}
	}
	i = len(triangle) - 1
	total = triangle[i][0]
	for j = 1; j < len(triangle[i]); j++ {
		if triangle[i][j] < total {
			total = triangle[i][j]
		}
	}
	return total
}
