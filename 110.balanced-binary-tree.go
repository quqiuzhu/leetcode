package main

/*
 * @lc app=leetcode id=110 lang=golang
 *
 * [110] Balanced Binary Tree
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalancedHeight(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}
	left, balanced := isBalancedHeight(root.Left)
	if !balanced {
		return 0, false
	}
	right, balanced := isBalancedHeight(root.Right)
	if !balanced {
		return 0, false
	}
	height := left
	if right > left {
		height = right
	}
	if -1 <= right-left && right-left <= 1 {
		return height + 1, true
	}
	return 0, false
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	_, ok := isBalancedHeight(root)
	return ok
}
