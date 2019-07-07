package main

/*
 * @lc app=leetcode id=27 lang=golang
 *
 * [27] Remove Element
 */
func removeElement(nums []int, val int) int {
	i, j := 0, len(nums)-1
	for i <= j {
		if nums[i] == val {
			for j >= 0 && nums[j] == val {
				j--
			}
			if i < j && nums[j] != val {
				nums[i], nums[j] = nums[j], nums[i]
				i++
			}
		} else {
			i++
		}
	}
	return i
}
