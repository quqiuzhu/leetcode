package main

/*
 * @lc app=leetcode id=131 lang=golang
 *
 * [131] Palindrome Partitioning
 */
func isPalindromeIII(s string) bool {
	ok := true
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			ok = false
			break
		}
	}
	return ok
}

func partPalindrome(s string) [][]string {
	if len(s) == 0 {
		return nil
	}
	if len(s) == 1 {
		return [][]string{[]string{s}}
	}
	var results [][]string
	for i := 1; i < len(s); i++ {
		if !isPalindromeIII(s[len(s)-i:]) {
			continue
		}
		rs := partPalindrome(s[:len(s)-i])
		for _, r := range rs {
			r = append(r, s[len(s)-i:])
			results = append(results, r)
		}
	}
	if isPalindromeIII(s) {
		results = append(results, []string{s})
	}
	return results
}

//与86题重名， 提交时取消下列注释
// func partition(s string) [][]string {
// 	return partPalindrome(s)
// }
