package main

/*
 * @lc app=leetcode id=84 lang=golang
 *
 * [84] Largest Rectangle in Histogram
 */
func largestRectangleArea(heights []int) int {
	var max, area, i, j, h int
	for i = 0; i < len(heights); i++ {
		if i < len(heights)-1 && heights[i] < heights[i+1] {
			continue
		}
		h = heights[i]
		for j = i; j >= 0; j-- {
			if heights[j] < h {
				h = heights[j]
			}
			area = h * (i - j + 1)
			if area > max {
				max = area
			}
		}
	}
	return max
}
