package main

import (
	"sort"
)

/*
 * @lc app=leetcode id=15 lang=golang
 *
 * [15] 3Sum
 */
func threeSum(nums []int) [][]int {
	sort.Sort(sort.IntSlice(nums))
	// for i := 0; i < len(nums)-1; i++ {
	// 	fmt.Print(nums[i], " ")
	// }
	// fmt.Println()
	var results [][]int
	begin, endPosition := 0, len(nums)-1
	for begin < endPosition {
		mid := begin + 1
		end := endPosition
		for mid < end {
			sum := nums[begin] + nums[mid] + nums[end]
			if sum == 0 {
				results = append(results, []int{nums[begin], nums[mid], nums[end]})
			}
			if sum <= 0 {
				mid++
				for nums[mid] == nums[mid-1] && mid < end {
					mid++
				}
			}
			if sum >= 0 {
				end--
				for nums[end] == nums[end+1] && mid < end {
					end--
				}
			}
		}

		begin++
		for nums[begin] == nums[begin-1] && begin < end {
			begin++
		}
	}
	return results
}
