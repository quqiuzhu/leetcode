package main

/*
 * @lc app=leetcode id=130 lang=golang
 *
 * [130] Surrounded Regions
 */

type uf struct {
	board      [][]byte
	surrounded map[int]bool
	keys       map[int]int
}

func (u *uf) find(i, j int) (int, bool) {
	k := i*len(u.board[0]) + j
	v, ok := u.keys[k]
	if !ok {
		v = k
	}
	return v, ok
}

func (u *uf) isSurrounded(i, j int) bool {
	k, _ := u.find(i, j)
	return u.surrounded[k]
}

func (u *uf) tryUnion(i, j int) {
	if u.board[i][j] == 'X' {
		return
	}
	k, ok := u.find(i, j)
	if ok {
		return
	}
	u.keys[k] = k
	isSurrounded := i > 0 && i < len(u.board)-1 && j > 0 && j < len(u.board[0])-1
	u.surrounded[k] = isSurrounded
	u.union(k, i, j)
}

func (u *uf) union(v, i, j int) {
	var k, m int
	var ok bool
	k = i*len(u.board[0]) + j
	u.keys[k] = v
	isSurrounded := i > 0 && i < len(u.board)-1 && j > 0 && j < len(u.board[0])-1
	if !isSurrounded {
		u.surrounded[v] = false
	}
	try := [8]int{i - 1, j, i, j - 1, i + 1, j, i, j + 1}
	for m = 0; m < 4; m++ {
		i = try[2*m]
		j = try[2*m+1]
		if i >= 0 && i < len(u.board) && j >= 0 && j < len(u.board[0]) && u.board[i][j] != 'X' {
			if k, ok = u.find(i, j); !ok {
				u.union(v, i, j)
			}
		}
	}
}

func solve(board [][]byte) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}
	var i, j int
	u := &uf{
		board:      board,
		surrounded: make(map[int]bool),
		keys:       make(map[int]int, len(board)*len(board[0])),
	}
	for i = 0; i < len(board); i++ {
		for j = 0; j < len(board[0]); j++ {
			u.tryUnion(i, j)
		}
	}

	for i = 0; i < len(board); i++ {
		for j = 0; j < len(board[0]); j++ {
			if board[i][j] == 'X' {
				continue
			}
			if u.isSurrounded(i, j) {
				board[i][j] = 'X'
			}
		}
	}
}
