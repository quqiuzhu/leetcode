package main

/*
 * @lc app=leetcode id=125 lang=golang
 *
 * [125] Valid Palindrome
 */
func validChar(b byte) (byte, bool) {
	if 'A' <= b && b <= 'Z' {
		b = 'a' + (b - 'A')
	}
	if 'a' <= b && b <= 'z' || '0' <= b && b <= '9' {
		return b, true
	}
	return b, false
}

func isPalindromeII(s string) bool {
	var chars []byte
	var ok bool
	var b byte
	for i := 0; i < len(s); i++ {
		b, ok = validChar(s[i])
		if ok {
			chars = append(chars, b)
		}
	}
	ok = true
	l := len(chars)
	for i := 0; i < (l+1)/2; i++ {
		if chars[i] != chars[l-1-i] {
			ok = false
			break
		}
	}
	return ok
}

//与9题重名， 提交时取消下列注释
// func isPalindrome(s string) bool {
// 	return isPalindromeII(s)
// }
