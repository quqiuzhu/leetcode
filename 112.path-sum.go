package main

/*
 * @lc app=leetcode id=112 lang=golang
 *
 * [112] Path Sum
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	sum -= root.Val
	if root.Left == nil && root.Right == nil {
		return sum == 0
	}

	var ok bool
	ok = hasPathSum(root.Left, sum)
	if ok {
		return true
	}
	ok = hasPathSum(root.Right, sum)
	return ok
}
