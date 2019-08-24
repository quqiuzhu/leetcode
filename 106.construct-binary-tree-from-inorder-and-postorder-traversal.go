package main

/*
 * @lc app=leetcode id=106 lang=golang
 *
 * [106] Construct Binary Tree from Inorder and Postorder Traversal
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func indexForTreeII(x int, xs []int) int {
	for i := 0; i < len(xs); i++ {
		if xs[i] == x {
			return i
		}
	}
	return -1
}

func buildTreeII(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: postorder[len(postorder)-1]}
	rootIdx := indexForTreeII(root.Val, inorder)
	root.Left = buildTreeII(inorder[:rootIdx], postorder[:rootIdx])
	root.Right = buildTreeII(inorder[rootIdx+1:], postorder[rootIdx:len(postorder)-1])
	return root
}

//与105题重名， 提交时取消下列注释
// func buildTree(inorder []int, postorder []int) *TreeNode {
// 	return buildTreeII(inorder, postorder)
// }
