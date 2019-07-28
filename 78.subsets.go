package main

/*
 * @lc app=leetcode id=78 lang=golang
 *
 * [78] Subsets
 */

func combineNums(nums []int, k int) [][]int {
	if len(nums) < k {
		return nil
	}

	var r, rx [][]int
	if k == 1 {
		for i := 0; i < len(nums); i++ {
			r = append(r, []int{nums[i]})
		}
		return r
	}

	r = combineNums(nums[1:], k)
	rx = combineNums(nums[1:], k-1)
	for i := 0; i < len(rx); i++ {
		rx[i] = append(rx[i], nums[0])
	}
	r = append(r, rx...)
	return r
}

func subsets(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{nil}
	}

	var set, sub [][]int
	for i := 1; i <= len(nums); i++ {
		sub = combineNums(nums, i)
		set = append(set, sub...)
	}
	set = append(set, nil)
	return set
}
