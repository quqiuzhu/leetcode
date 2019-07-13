package main

import "sort"

/*
 * @lc app=leetcode id=40 lang=golang
 *
 * [40] Combination Sum II
 */

func recursiveCombinationSum2(candidates []int, target int) [][]int {
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
	for i = lc - 1; i >= 0; {
		if target == candidates[i] {
			x = append(x, []int{candidates[i]})
		} else {
			r = recursiveCombinationSum2(candidates[:i], target-candidates[i])
			if len(r) > 0 {
				for j = 0; j < len(r); j++ {
					r[j] = append(r[j], candidates[i])
					x = append(x, r[j])
				}
			}
		}
		i--
		for i >= 0 && candidates[i] == candidates[i+1] {
			i--
		}
	}
	return x
}

func combinationSum2(candidates []int, target int) [][]int {
	sort.Sort(sort.IntSlice(candidates))
	return recursiveCombinationSum2(candidates, target)
}
