package main

/*
 * @lc app=leetcode id=113 lang=golang
 *
 * [113] Path Sum II
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, sum int) [][]int {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		if root.Val == sum {
			return [][]int{
				[]int{root.Val},
			}
		}
		return nil
	}
	sum -= root.Val
	paths := pathSum(root.Left, sum)
	paths = append(paths, pathSum(root.Right, sum)...)
	for i := range paths {
		paths[i] = append([]int{root.Val}, paths[i]...)
	}
	return paths
}
