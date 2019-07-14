package main

/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */
func twoSum(nums []int, target int) []int {
	var i, j, t int
	var ok bool
	index := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		index[nums[i]] = i
	}

	for i = 0; i < len(nums); i++ {
		t = target - nums[i]
		if j, ok = index[t]; ok && j != i {
			break
		}
	}

	return []int{i, j}
}
