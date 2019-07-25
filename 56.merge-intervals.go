package main

import "sort"

/*
 * @lc app=leetcode id=56 lang=golang
 *
 * [56] Merge Intervals
 */
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	var rs [][]int
	var item []int
	for i := 0; i < len(intervals); i++ {
		if item == nil {
			item = intervals[i]
		} else if item[1] < intervals[i][0] {
			rs = append(rs, item)
			item = intervals[i]
		} else if item[1] < intervals[i][1] {
			item[1] = intervals[i][1]
		}
	}
	if len(item) != 0 {
		rs = append(rs, item)
	}
	return rs
}
