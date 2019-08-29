package main

/*
 * @lc app=leetcode id=129 lang=golang
 *
 * [129] Sum Root to Leaf Numbers
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return root.Val
	}
	var sum int
	if root.Left != nil {
		root.Left.Val += root.Val * 10
		sum += sumNumbers(root.Left)
	}
	if root.Right != nil {
		root.Right.Val += root.Val * 10
		sum += sumNumbers(root.Right)
	}
	return sum
}
