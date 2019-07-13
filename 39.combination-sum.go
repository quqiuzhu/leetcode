package main

import "sort"

/*
 * @lc app=leetcode id=39 lang=golang
 *
 * [39] Combination Sum
 */

func recursiveCombinationSum(candidates []int, target int) [][]int {
	lc := len(candidates)
	if lc == 0 {
		return nil
	}
	least := candidates[0]
	if target < least {
		return nil
	}

	var i, j int
	var r, x [][]int
	for i = lc - 1; i >= 0; i-- {
		if target == candidates[i] {
			x = append(x, []int{candidates[i]})
		} else {
			r = recursiveCombinationSum(candidates[:i+1], target-candidates[i])
			if len(r) > 0 {
				for j = 0; j < len(r); j++ {
					r[j] = append(r[j], candidates[i])
					x = append(x, r[j])
				}
			}
		}
	}
	return x
}

func combinationSum(candidates []int, target int) [][]int {
	sort.Sort(sort.IntSlice(candidates))
	return recursiveCombinationSum(candidates, target)
}
