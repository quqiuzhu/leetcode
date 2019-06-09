package main

import "fmt"

/*
 * @lc app=leetcode id=9 lang=golang
 *
 * [9] Palindrome Number
 */
func isPalindromeNonStringMethod(x int) bool {
	if x < 0 {
		return false
	} else if x == 0 {
		return true
	}

	originalVal := x
	newValue := 0
	for x != 0 {
		newValue = newValue*10 + x%10
		x = x / 10
	}

	return originalVal == newValue
}

func isPalindromeStringMethod(x int) bool {
	s := fmt.Sprintf("%d", x)
	l := len(s)
	// fmt.Println("s:", s, "l:", l)
	for i := 0; i < l/2+1; i++ {
		if s[i] != s[l-i-1] {
			return false
		}
	}
	return true
}

func isPalindrome(x int) bool {
	return isPalindromeNonStringMethod(x)
}
