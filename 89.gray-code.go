package main

/*
 * @lc app=leetcode id=89 lang=golang
 *
 * [89] Gray Code
 */
func grayCode(n int) []int {
	if n == 0 {
		return []int{0}
	}
	if n == 1 {
		return []int{0, 1}
	}

	r := grayCode(n - 1)
	l := len(r)
	x := 1 << uint(n-1)
	for i := 0; i < l; i++ {
		r = append(r, r[l-1-i]+x)
	}
	return r
}
