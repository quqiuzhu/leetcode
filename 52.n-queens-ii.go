package main

/*
 * @lc app=leetcode id=52 lang=golang
 *
 * [52] N-Queens II
 */
func isValidBoardII(board []int, r, c int) bool {
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

func dfsII(count *int, board []int, row int) {
	if row == len(board) {
		*count = *count + 1
		return
	}

	for i := 0; i < len(board); i++ {
		board[row] = i
		if isValidBoardII(board, row, i) {
			dfsII(count, board, row+1)
		}
		board[row] = -1
	}
}

func totalNQueens(n int) int {
	var count int
	board := make([]int, n)
	for i := 0; i < len(board); i++ {
		board[i] = -1
	}
	dfsII(&count, board, 0)
	return count
}
