package main

/*
 * @lc app=leetcode id=114 lang=golang
 *
 * [114] Flatten Binary Tree to Linked List
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	if root.Left == nil {
		flatten(root.Right)
		return
	}

	it := root.Left
	for it.Right != nil {
		it = it.Right
	}
	it.Right = root.Right
	root.Right = root.Left
	root.Left = nil
	flatten(root.Right)
}
