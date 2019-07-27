package main

/*
 * @lc app=leetcode id=58 lang=golang
 *
 * [58] Length of Last Word
 */
func lengthOfLastWord(s string) int {
	var length int
	found := false
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if found {
				break
			}
			continue
		}
		found = true
		length++
	}
	return length
}
