package main

/*
 * @lc app=leetcode id=46 lang=golang
 *
 * [46] Permutations
 */
func permute(nums []int) [][]int {
	if len(nums) == 1 {
		n := make([]int, len(nums))
		copy(n, nums)
		return [][]int{n}
	}

	var i, j int
	var r, x [][]int

	for i = 0; i < len(nums); i++ {
		n := make([]int, len(nums))
		copy(n, nums)
		n[0], n[i] = n[i], n[0]
		x = permute(n[1:])
		for j = 0; j < len(x); j++ {
			x[j] = append(x[j], n[0])
		}
		r = append(r, x...)
	}
	return r
}
