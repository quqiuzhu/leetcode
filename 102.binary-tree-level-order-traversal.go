package main

/*
 * @lc app=leetcode id=102 lang=golang
 *
 * [102] Binary Tree Level Order Traversal
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	var cur, next []*TreeNode
	var r [][]int
	cur = append(cur, root)
	for len(cur) > 0 {
		var level []int
		for i := 0; i < len(cur); i++ {
			if cur[i] != nil {
				level = append(level, cur[i].Val)
				next = append(next, cur[i].Left)
				next = append(next, cur[i].Right)
			}
		}
		if level != nil {
			r = append(r, level)
		}
		cur = next
		next = nil
	}
	return r
}
