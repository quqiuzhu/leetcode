package main

/*
 * @lc app=leetcode id=95 lang=golang
 *
 * [95] Unique Binary Search Trees II
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func generateTree(begin, end int) []*TreeNode {
	if begin > end {
		return []*TreeNode{nil}
	}
	if begin == end {
		return []*TreeNode{&TreeNode{
			Val: begin,
		}}
	}
	var trees []*TreeNode
	for i := begin; i <= end; i++ {
		lefts := generateTree(begin, i-1)
		rights := generateTree(i+1, end)
		for _, left := range lefts {
			for _, right := range rights {
				root := &TreeNode{
					Val:   i,
					Left:  left,
					Right: right,
				}
				trees = append(trees, root)
			}
		}
	}
	return trees
}

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	return generateTree(1, n)
}
