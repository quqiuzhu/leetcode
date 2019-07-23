package main

/*
 * @lc app=leetcode id=51 lang=golang
 *
 * [51] N-Queens
 */

func isValidBoard(board []int, r, c int) bool {
	for i := 0; i < r; i++ {
		if board[i] == board[r] {
			return false
		}

		v := board[i] - board[r]
		if v < 0 {
			v = -v
		}
		if r-i == v {
			return false
		}
	}
	return true
}

func dfs(rs *[][]string, board []int, row int) {
	if row == len(board) {
		var r []string
		b := make([]byte, len(board))
		for i := 0; i < len(board); i++ {
			for j := 0; j < len(board); j++ {
				b[j] = '.'
				if j == board[i] {
					b[j] = 'Q'
				}
			}
			r = append(r, string(b))
		}

		*rs = append(*rs, r)
		return
	}

	for i := 0; i < len(board); i++ {
		board[row] = i
		if isValidBoard(board, row, i) {
			dfs(rs, board, row+1)
		}
		board[row] = -1
	}
}

func solveNQueens(n int) [][]string {
	rs := make([][]string, 0)
	board := make([]int, n)
	for i := 0; i < len(board); i++ {
		board[i] = -1
	}
	dfs(&rs, board, 0)
	return rs
}
