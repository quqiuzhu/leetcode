package main

/*
 * @lc app=leetcode id=139 lang=golang
 *
 * [139] Word Break
 */
func wordBreak(s string, wordDict []string) bool {
	var i, j int
	sets := make(map[string]bool, len(wordDict))
	for i = 0; i < len(wordDict); i++ {
		sets[wordDict[i]] = true
	}
	results := make([]bool, len(s)+1)
	results[0] = true
	for i = 1; i <= len(s); i++ {
		for j = i - 1; j >= 0; j-- {
			if results[j] && sets[s[j:i]] {
				results[i] = true
				break
			}
		}
	}
	return results[len(s)]
}
