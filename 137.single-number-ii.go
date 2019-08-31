package main

/*
 * @lc app=leetcode id=137 lang=golang
 *
 * [137] Single Number II
 * 位操作太过 tricky
 */
func singleNumberII(nums []int) int {
	once, twice, triple := 0, 0, 0
	for i := 0; i < len(nums); i++ {
		twice |= (once & nums[i])
		once = once ^ nums[i]
		triple = ^(once & twice)
		once &= triple
		twice &= triple
	}
	return once
}

// 与136题重名， 提交时取消下列注释
// func singleNumber(nums []int) int {
// 	return singleNumberII(nums)
// }
