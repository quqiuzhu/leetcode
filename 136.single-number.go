package main

/*
 * @lc app=leetcode id=136 lang=golang
 *
 * [136] Single Number
 */
func singleNumber(nums []int) int {
	v := 0
	for _, num := range nums {
		v ^= num
	}
	return v
}
