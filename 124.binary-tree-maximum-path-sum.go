package main

/*
 * @lc app=leetcode id=124 lang=golang
 *
 * [124] Binary Tree Maximum Path Sum
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSums(root *TreeNode) (int, int) {
	if root.Left == nil && root.Right == nil {
		return root.Val, root.Val
	}
	var max, maxOverRoot int
	if root.Left == nil || root.Right == nil {
		child := root.Left
		if child == nil {
			child = root.Right
		}
		max, maxOverRoot = maxPathSums(child)
		if max > 0 {
			max += root.Val
		} else {
			max = root.Val
		}

		if max > maxOverRoot {
			maxOverRoot = max
		}
		return max, maxOverRoot
	}

	left, leftOverRoot := maxPathSums(root.Left)
	right, rightOverRoot := maxPathSums(root.Right)
	max = left
	if max < right {
		max = right
	}
	if max < 0 {
		max = 0
	}
	max += root.Val

	maxOverRoot = left + right + root.Val
	if leftOverRoot > maxOverRoot {
		maxOverRoot = leftOverRoot
	}
	if rightOverRoot > maxOverRoot {
		maxOverRoot = rightOverRoot
	}
	if maxOverRoot < max {
		maxOverRoot = max
	}
	return max, maxOverRoot
}

func maxPathSum(root *TreeNode) int {
	sum, sumOverRoot := maxPathSums(root)
	if sumOverRoot > sum {
		sum = sumOverRoot
	}
	return sum
}
