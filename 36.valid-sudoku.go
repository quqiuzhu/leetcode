package main

/*
 * @lc app=leetcode id=36 lang=golang
 *
 * [36] Valid Sudoku
 */
func isValidSudoku(board [][]byte) bool {
	var i, j, k, l int
	var box map[byte]int
	var c byte
	column := make([]map[byte]int, 9)
	row := make([]map[byte]int, 9)
	for i = 0; i < 9; i++ {
		column[i] = make(map[byte]int)
		row[i] = make(map[byte]int)
	}

	for i = 0; i < 3; i++ {
		for j = 0; j < 3; j++ {
			box = make(map[byte]int)
			for k = 0; k < 3; k++ {
				for l = 0; l < 3; l++ {
					c = board[i*3+k][j*3+l]
					if c == '.' {
						continue
					}

					column[i*3+k][c] = column[i*3+k][c] + 1
					if column[i*3+k][c] > 1 {
						return false
					}

					row[j*3+l][c] = row[j*3+l][c] + 1
					if row[j*3+l][c] > 1 {
						return false
					}

					box[c] = box[c] + 1
					if box[c] > 1 {
						return false
					}
				}
			}
		}
	}
	return true
}
