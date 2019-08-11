package main

import "sort"

/*
 * @lc app=leetcode id=90 lang=golang
 *
 * [90] Subsets II
 */
func subsetsWithDup(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{nil}
	}

	sort.Sort(sort.IntSlice(nums))

	var result [][]int
	result = append(result, nil)
	var i, j, size, l int
	for i = 0; i < len(nums); i++ {
		l = len(result)
		for j = 0; j < l; j++ {
			if i == 0 || nums[i] != nums[i-1] || j >= size {
				r := make([]int, len(result[j])+1)
				copy(r, result[j])
				r[len(r)-1] = nums[i]
				result = append(result, r)
			}
		}
		size = l
	}
	return result
}
