package main

import "sort"

/*
 * @lc app=leetcode id=16 lang=golang
 *
 * [16] 3Sum Closest
 */

func diff(a, b int) int {
	v := a - b
	if v < 0 {
		v = -v
	}
	return v
}

func threeSumClosest(nums []int, target int) int {
	sort.Sort(sort.IntSlice(nums))
	begin, endPosition := 0, len(nums)-1

	value := nums[0] + nums[1] + nums[2]
	closest := diff(value, target)

	for begin < endPosition {
		mid := begin + 1
		end := endPosition
		for mid < end {
			sum := nums[begin] + nums[mid] + nums[end]
			d := diff(sum, target)
			if d < closest {
				closest = d
				value = sum
			}
			if sum == target {
				return target
			}
			if sum < target {
				mid++
				for nums[mid] == nums[mid-1] && mid < end {
					mid++
				}
			}
			if sum > target {
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
	return value
}
