package main

/*
 * @lc app=leetcode id=44 lang=golang
 *
 * [44] Wildcard Matching
 */

func dpIsMatch(s string, p string) bool {
	ls, lp := len(s)+1, len(p)+1

	match := make([][]bool, lp)
	var i, j, k int
	for i = 0; i < lp; i++ {
		match[i] = make([]bool, ls)
	}
	match[0][0] = true

	for i = 1; i < lp; i++ {
		switch p[i-1] {
		case '*':
			for k = 0; k < ls; k++ {
				if match[i-1][k] {
					break
				}
			}
			for j = k; j < ls; j++ {
				match[i][j] = true
			}

		default:
			for j = 0; j < ls-1; j++ {
				if match[i-1][j] {
					match[i][j+1] = p[i-1] == s[j] || p[i-1] == '?'
				}
			}
		}
	}

	return match[lp-1][ls-1]
}

// 与第十题函数名冲突， 提交时取消下列注释
// func isMatch(s string, p string) bool {
// 	return dpIsMatch(s, p)
// }
