package main

/*
 * @lc app=leetcode id=118 lang=golang
 *
 * [118] Pascal's Triangle
 */
func generate(numRows int) [][]int {
	if numRows == 0 {
		return nil
	}
	var i, j, val int
	var r [][]int
	r = [][]int{[]int{1}}
	for i = 1; i < numRows; i++ {
		line := make([]int, i+1)
		for j = 0; j < i+1; j++ {
			val = 1
			if j > 0 && j < i {
				val = r[i-1][j-1] + r[i-1][j]
			}
			line[j] = val
		}
		r = append(r, line)
	}
	return r
}
