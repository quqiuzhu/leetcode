package main

/*
 * @lc app=leetcode id=37 lang=golang
 *
 * [37] Sudoku Solver
 */

func recursiveSolveSudoku(board [][]byte, c, r, b []map[byte]int) bool {
	var i, j, k, l int
	var cc byte

	for i = 0; i < 3; i++ {
		for j = 0; j < 3; j++ {
			for k = 0; k < 3; k++ {
				for l = 0; l < 3; l++ {
					cc = board[i*3+k][j*3+l]
					if cc != '.' {
						continue
					}

					for cc = '1'; cc <= '9'; cc++ {
						if c[i*3+k][cc] == 0 && r[j*3+l][cc] == 0 && b[i*3+j][cc] == 0 {
							board[i*3+k][j*3+l] = cc
							c[i*3+k][cc] = 1
							r[j*3+l][cc] = 1
							b[i*3+j][cc] = 1
							if recursiveSolveSudoku(board, c, r, b) {
								return true
							}
							c[i*3+k][cc] = 0
							r[j*3+l][cc] = 0
							b[i*3+j][cc] = 0
							board[i*3+k][j*3+l] = '.'
						}
					}

					return false
				}
			}
		}
	}

	return true
}

func solveSudoku(board [][]byte) {
	var i, j, k, l int
	var c byte

	column := make([]map[byte]int, 9)
	row := make([]map[byte]int, 9)
	box := make([]map[byte]int, 9)
	for i = 0; i < 9; i++ {
		column[i] = make(map[byte]int)
		row[i] = make(map[byte]int)
		box[i] = make(map[byte]int)
	}

	for i = 0; i < 3; i++ {
		for j = 0; j < 3; j++ {
			for k = 0; k < 3; k++ {
				for l = 0; l < 3; l++ {
					c = board[i*3+k][j*3+l]
					if c == '.' {
						continue
					}
					column[i*3+k][c] = 1
					row[j*3+l][c] = 1
					box[i*3+j][c] = 1
				}
			}
		}
	}
	recursiveSolveSudoku(board, column, row, box)
}
