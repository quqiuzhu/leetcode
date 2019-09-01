package main

/*
 * @lc app=leetcode id=144 lang=golang
 *
 * [144] Binary Tree Preorder Traversal
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var r []int
	r = append(r, root.Val)
	r = append(r, preorderTraversal(root.Left)...)
	r = append(r, preorderTraversal(root.Right)...)
	return r
}
