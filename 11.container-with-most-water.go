package main

/*
 * @lc app=leetcode id=11 lang=golang
 *
 * [11] Container With Most Water
 */

func forceMaxArea(height []int) int {
	length := len(height)
	max := 0
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			h := height[i]
			if height[j] < h {
				h = height[j]
			}
			area := (j - i) * h
			if area > max {
				max = area
			}
		}
	}
	return max
}

func walkMaxArea(height []int) int {
	begin, end := 0, len(height)-1
	max := 0
	for begin < end {
		h := height[begin]
		if height[end] < h {
			h = height[end]
		}
		area := (end - begin) * h
		if area > max {
			max = area
		}

		if height[begin] <= height[end] {
			begin++
		} else {
			end--
		}
	}
	return max
}

func maxArea(height []int) int {
	return walkMaxArea(height)
}
