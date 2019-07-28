package main

/*
 * @lc app=leetcode id=77 lang=golang
 *
 * [77] Combinations
 */
func combine(n int, k int) [][]int {
	if n < k {
		return nil
	}

	var r, rx [][]int
	if k == 1 {
		for i := 1; i <= n; i++ {
			r = append(r, []int{i})
		}
		return r
	}

	r = combine(n-1, k)
	rx = combine(n-1, k-1)
	for i := 0; i < len(rx); i++ {
		rx[i] = append(rx[i], n)
	}
	r = append(r, rx...)
	return r
}
