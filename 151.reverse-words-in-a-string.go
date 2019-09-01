package main

/*
 * @lc app=leetcode id=151 lang=golang
 *
 * [151] Reverse Words in a String
 */
func reverseWords(s string) string {
	var words [][]byte
	var i, j int
	for i = 0; i < len(s); i++ {
		if s[i] == ' ' {
			if j < i {
				words = append(words, []byte(s[j:i]))
			}
			j = i + 1
		} else if i == len(s)-1 && j <= i {
			words = append(words, []byte(s[j:]))
		}
	}

	var result []byte
	for i = len(words) - 1; i >= 0; i-- {
		result = append(result, words[i]...)
		if i > 0 {
			result = append(result, ' ')
		}
	}
	return string(result)
}
