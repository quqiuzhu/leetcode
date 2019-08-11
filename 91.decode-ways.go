package main

/*
 * @lc app=leetcode id=91 lang=golang
 *
 * [91] Decode Ways
 */
func numDecodings(s string) int {
	if len(s) == 0 || s[0] == '0' {
		return 0
	}
	if len(s) == 1 {
		return 1
	}

	if s[1] == '0' && s[0] > '2' {
		return 0
	}

	counts := make([]int, len(s))
	counts[0] = 1
	counts[1] = 1
	if (s[0] == '1' || s[0] == '2' && s[1] <= '6') && s[1] != '0' {
		counts[1] = 2
	}
	for i := 2; i < len(s); i++ {
		if s[i] == '0' && (s[i-1] == '0' || s[i-1] > '2') {
			break
		}
		if s[i] == '0' {
			counts[i] = counts[i-2]
			continue
		}
		counts[i] = counts[i-1]
		if s[i-1] == '1' || s[i-1] == '2' && s[i] <= '6' {
			counts[i] += counts[i-2]
		}
	}
	return counts[len(s)-1]
}
