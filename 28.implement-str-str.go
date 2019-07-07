package main

/*
 * @lc app=leetcode id=28 lang=golang
 *
 * [28] Implement strStr()
 */
func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	var i, j int
	var found bool
	for i = 0; i <= len(haystack)-len(needle); i++ {
		found = true
		for j = 0; j < len(needle); j++ {
			if haystack[i+j] != needle[j] {
				found = false
				break
			}
		}
		if found {
			return i
		}
	}

	return -1
}
