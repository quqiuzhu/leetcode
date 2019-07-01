package main

/*
 * @lc app=leetcode id=14 lang=golang
 *
 * [14] Longest Common Prefix
 */

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	var prefix []byte
	for j := 0; ; j++ {
		if len(strs[0]) <= j {
			break
		}

		common, ch := true, strs[0][j]
		for i := 1; i < len(strs); i++ {
			if len(strs[i]) <= j || strs[i][j] != ch {
				common = false
				break
			}
		}

		if common {
			prefix = append(prefix, ch)
		} else {
			break
		}
	}
	return string(prefix)
}
