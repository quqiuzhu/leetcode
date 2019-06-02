package main

import "fmt"

/*
 * @lc app=leetcode id=5 lang=golang
 *
 * [5] Longest Palindromic Substring
 */

// l length of s
// index max expanded substring center index
// max max expanded substring single side length
//
func longestPalindrome(s string) string {
	index, max, cur := 0, 0, 0
	l, lx := len(s), 2*len(s)+1
	// fmt.Println("start", l, lx)
	for i := 0; i < lx; i++ {
		cur = 0
		for j := 0; j <= l && i-j >= 0 && i+j < lx; j++ {
			// fmt.Print(i, j, " ")
			// jump #
			if (i+j)%2 == 0 {
				cur = j
				continue
			}
			// fmt.Println(s[(i+j-1)/2], s[(i-j-1)/2])
			if s[(i+j-1)/2] == s[(i-j-1)/2] {
				cur = j
			} else {
				break
			}
		}
		if cur > max {
			// fmt.Println("update", i, cur)
			index = i
			max = cur
		}
	}
	i, j := index, max
	// fmt.Println("stop", i, j)
	// slice char s[(i-j-1)/2] to s[(i+j-1)/2]
	return s[(i-j-1)/2 : (i+j+1)/2]
}

func main() {
	fmt.Println(longestPalindrome("bababd"))
}
