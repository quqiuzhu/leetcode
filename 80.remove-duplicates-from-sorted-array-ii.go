package main

/*
 * @lc app=leetcode id=80 lang=golang
 *
 * [80] Remove Duplicates from Sorted Array II
 */

func removeDuplicatesII(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	var i, j, item, count int
	item = nums[0]
	count = 1
	i = 1
	for j = 1; j < len(nums); j++ {
		if nums[j] == item {
			count++
		} else {
			count = 1
		}
		item = nums[j]

		if count <= 2 {
			nums[i] = nums[j]
			i++
		}
	}
	return i
}

// 与26题重名， 提交时取消注释
// func removeDuplicates(nums []int) int {
// 	return removeDuplicatesII(nums)
// }
