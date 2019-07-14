package main

import "sort"

/*
 * @lc app=leetcode id=47 lang=golang
 *
 * [47] Permutations II
 */

func permuteUnique(nums []int) [][]int {
	if len(nums) == 1 {
		n := make([]int, len(nums))
		copy(n, nums)
		return [][]int{n}
	}

	sort.Sort(sort.IntSlice(nums))

	var i, j int
	var r, x [][]int

	for i = 0; i < len(nums); {
		n := make([]int, len(nums))
		copy(n, nums)
		n[0], n[i] = n[i], n[0]
		x = permuteUnique(n[1:])
		for j = 0; j < len(x); j++ {
			x[j] = append(x[j], n[0])
		}
		r = append(r, x...)
		i++
		for i < len(nums) && nums[i-1] == nums[i] {
			i++
		}
	}
	return r
}
