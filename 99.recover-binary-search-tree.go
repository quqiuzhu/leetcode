package main

/*
 * @lc app=leetcode id=99 lang=golang
 *
 * [99] Recover Binary Search Tree
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderNodes(root *TreeNode) []*TreeNode {
	if root == nil {
		return nil
	}
	var r []*TreeNode
	lefts := inorderNodes(root.Left)
	rights := inorderNodes(root.Right)
	r = append(lefts, root)
	r = append(r, rights...)
	return r
}

func recoverTree(root *TreeNode) {
	nodes := inorderNodes(root)
	var i, j int
	for i = 0; i < len(nodes)-1; i++ {
		if nodes[i].Val >= nodes[i+1].Val {
			break
		}
	}
	for j = len(nodes) - 1; j > 0; j-- {
		if nodes[j].Val <= nodes[j-1].Val {
			break
		}
	}
	nodes[i].Val, nodes[j].Val = nodes[j].Val, nodes[i].Val
}
