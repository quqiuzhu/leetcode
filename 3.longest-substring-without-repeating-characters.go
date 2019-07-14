package main

/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 */
func lengthOfLongestSubstring(s string) int {
	index := make(map[byte]int)
	var i, j, v, max, c int
	var ok bool

	for i = 0; i < len(s); i++ {
		v, ok = index[s[i]]
		if ok {
			c = i - v - 1
			for ; j <= v; j++ {
				delete(index, s[j])
			}
		}

		index[s[i]] = i
		c++
		if c > max {
			max = c
		}
	}

	return max
}
