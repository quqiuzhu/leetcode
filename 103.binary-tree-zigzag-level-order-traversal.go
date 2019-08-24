package main

/*
 * @lc app=leetcode id=103 lang=golang
 *
 * [103] Binary Tree Zigzag Level Order Traversal
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
	var cur, next []*TreeNode
	var r [][]int
	var i, j int
	zig := true
	cur = append(cur, root)
	for len(cur) > 0 {
		var level []int
		for i = 0; i < len(cur); i++ {
			j = i
			if !zig {
				j = len(cur) - 1 - i
			}
			if cur[j] != nil {
				level = append(level, cur[j].Val)
			}
			if cur[i] != nil {
				next = append(next, cur[i].Left)
				next = append(next, cur[i].Right)
			}
		}

		if level != nil {
			r = append(r, level)
		}
		zig = !zig
		cur = next
		next = nil
	}
	return r
}
