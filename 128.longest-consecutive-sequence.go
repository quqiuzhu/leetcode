package main

/*
 * @lc app=leetcode id=128 lang=golang
 *
 * [128] Longest Consecutive Sequence
 */
func longestConsecutive(nums []int) int {
	m := make(map[int]bool)
	var ok bool
	var i, max, now int
	for i = 0; i < len(nums); i++ {
		m[nums[i]] = true
	}
	for i = range m {
		if _, ok = m[i-1]; !ok {
			now = 1
			for m[i+1] {
				now++
				i++
			}
			if now > max {
				max = now
			}
		}
	}
	return max
}
