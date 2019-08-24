package main

/*
 * @lc app=leetcode id=101 lang=golang
 *
 * [101] Symmetric Tree
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// Copy from 100, 提交时取消下列注释
// func isSameTree(p *TreeNode, q *TreeNode) bool {
// 	if p == nil && q == nil {
// 		return true
// 	} else if p == nil || q == nil {
// 		return false
// 	}

// 	if p.Val != q.Val {
// 		return false
// 	}

// 	if !isSameTree(p.Left, q.Left) {
// 		return false
// 	}

// 	if !isSameTree(p.Right, q.Right) {
// 		return false
// 	}
// 	return true
// }

func symmetricTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	newRoot := &TreeNode{
		Val:   root.Val,
		Left:  symmetricTree(root.Right),
		Right: symmetricTree(root.Left),
	}
	return newRoot
}

func isSymmetric(root *TreeNode) bool {
	symmetric := symmetricTree(root)
	return isSameTree(root, symmetric)
}
