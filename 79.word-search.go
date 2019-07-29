package main

/*
 * @lc app=leetcode id=79 lang=golang
 *
 * [79] Word Search
 */

func findFrom(board [][]byte, i, j int, b [][]bool, word string) bool {
	m, n := len(board), len(board[0])
	if len(word) == 0 {
		return true
	}
	if i < 0 || i >= m {
		return false
	}
	if j < 0 || j >= n {
		return false
	}
	if board[i][j] != word[0] {
		return false
	}
	if b[i][j] {
		return false
	}

	b[i][j] = true
	if findFrom(board, i-1, j, b, word[1:]) {
		return true
	}
	if findFrom(board, i, j-1, b, word[1:]) {
		return true
	}
	if findFrom(board, i+1, j, b, word[1:]) {
		return true
	}
	if findFrom(board, i, j+1, b, word[1:]) {
		return true
	}
	b[i][j] = false
	return false
}

func exist(board [][]byte, word string) bool {
	if len(board) == 0 {
		return false
	}
	m, n := len(board), len(board[0])
	var i, j int
	b := make([][]bool, m)
	for i = 0; i < m; i++ {
		b[i] = make([]bool, n)
	}
	for i = 0; i < m; i++ {
		for j = 0; j < n; j++ {
			if findFrom(board, i, j, b, word) {
				return true
			}
		}
	}
	return false
}
