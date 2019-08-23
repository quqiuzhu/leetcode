package main

/*
 * @lc app=leetcode id=98 lang=golang
 *
 * [98] Validate Binary Search Tree
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var it *TreeNode
	if root.Left != nil {
		it = root.Left
		for it.Right != nil {
			it = it.Right
		}
		if it.Val >= root.Val {
			return false
		}
		if !isValidBST(root.Left) {
			return false
		}
	}
	if root.Right != nil {
		it = root.Right
		for it.Left != nil {
			it = it.Left
		}
		if it.Val <= root.Val {
			return false
		}
		if !isValidBST(root.Right) {
			return false
		}
	}
	return true
}
