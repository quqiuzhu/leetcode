package main

/*
 * @lc app=leetcode id=13 lang=golang
 *
 * [13] Roman to Integer
 */
func romanToInt(s string) int {
	v := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	result, pre := 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		cur := v[s[i]]
		if cur < pre {
			result -= cur
		} else {
			result += cur
		}
		pre = cur
	}
	return result
}
