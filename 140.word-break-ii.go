package main

/*
 * @lc app=leetcode id=140 lang=golang
 *
 * [140] Word Break II
 */
func wordBreakIIResults(s string, m [][]int, i int) [][]byte {
	var results [][]byte
	var j, k, n int
	for n = 0; n < len(m[i]); n++ {
		j = m[i][n]
		if j == 0 {
			results = append(results, []byte(s[j:i]))
		} else {
			rs := wordBreakIIResults(s, m, j)
			for k = range rs {
				rs[k] = append(rs[k], ' ')
				rs[k] = append(rs[k], []byte(s[j:i])...)
				results = append(results, rs[k])
			}
		}
	}
	return results
}

func wordBreakII(s string, wordDict []string) []string {
	var i, j int
	set := make(map[string]bool, len(wordDict))
	for i = 0; i < len(wordDict); i++ {
		set[wordDict[i]] = true
	}
	reach := make([]bool, len(s)+1)
	m := make([][]int, len(s)+1)
	reach[0] = true
	for i = 1; i <= len(s); i++ {
		for j = i - 1; j >= 0; j-- {
			if reach[j] && set[s[j:i]] {
				reach[i] = true
				m[i] = append(m[i], j)
			}
		}
	}
	var results []string
	rs := wordBreakIIResults(s, m, len(s))
	for i = 0; i < len(rs); i++ {
		results = append(results, string(rs[i]))
	}
	return results
}

//与139题重名， 提交时取消下列注释
// func wordBreak(s string, wordDict []string) []string {
// 	return wordBreakII(s, wordDict)
// }
