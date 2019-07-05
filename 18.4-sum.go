package main

import "sort"

/*
 * @lc app=leetcode id=18 lang=golang
 *
 * [18] 4Sum
 */

type result struct {
	x1, x2, x3, x4 int
}

type sortBy []*result

func (a sortBy) Len() int      { return len(a) }
func (a sortBy) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a sortBy) Less(i, j int) bool {
	if a[i].x1 != a[j].x1 {
		return a[i].x1 < a[j].x1
	}
	if a[i].x2 != a[j].x2 {
		return a[i].x2 < a[j].x2
	}
	if a[i].x3 != a[j].x3 {
		return a[i].x3 < a[j].x3
	}
	return a[i].x4 <= a[j].x4
}

func fourSum(nums []int, target int) [][]int {
	sort.Sort(sort.IntSlice(nums))
	x1, x2, x3, x4, twoSum, end := 0, 0, 0, 0, 0, len(nums)-1
	twoSums := make(map[int][]*result)

	for x1 < end {
		x2 = x1 + 1
		for x2 < end {
			twoSum = nums[x1] + nums[x2]
			twoSums[twoSum] = append(twoSums[twoSum], &result{x1: x1, x2: x2})
			x2++
		}
		x1++
	}

	var rs []*result
	for x3 < end {
		x4 = x3 + 1
		for x4 <= end {
			twoSum = target - nums[x3] - nums[x4]
			if xs, ok := twoSums[twoSum]; ok {
				for _, x := range xs {
					if x.x2 < x3 {
						rs = append(rs, &result{
							x1: nums[x.x1],
							x2: nums[x.x2],
							x3: nums[x3],
							x4: nums[x4],
						})
					}
				}
			}
			x4++
		}
		x3++
	}

	sort.Sort(sortBy(rs))
	var r [][]int
	for i := 0; i < len(rs); i++ {
		if 0 < i && *rs[i] == *rs[i-1] {
			continue
		}
		r = append(r, []int{rs[i].x1, rs[i].x2, rs[i].x3, rs[i].x4})
	}
	return r
}
