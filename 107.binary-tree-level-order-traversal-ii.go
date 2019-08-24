package main

/*
 * @lc app=leetcode id=107 lang=golang
 *
 * [107] Binary Tree Level Order Traversal II
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottoms(roots []*TreeNode) [][]int {
	var next []*TreeNode
	var level []int
	for i := 0; i < len(roots); i++ {
		if roots[i] != nil {
			level = append(level, roots[i].Val)
			next = append(next, roots[i].Left)
			next = append(next, roots[i].Right)
		}
	}

	var levels [][]int
	if next != nil {
		levels = levelOrderBottoms(next)
	}
	if level != nil {
		levels = append(levels, level)
	}
	return levels
}

func levelOrderBottom(root *TreeNode) [][]int {
	return levelOrderBottoms([]*TreeNode{root})
}
