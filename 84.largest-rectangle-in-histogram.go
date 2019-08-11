package main

/*
 * @lc app=leetcode id=84 lang=golang
 *
 * [84] Largest Rectangle in Histogram
 */

func largestRectangleAreaStackMethod(heights []int) int {
	heights = append(heights, 0)
	var max, area, i, top, h, w int
	var stack []int
	for i < len(heights) {
		top = len(stack) - 1
		if len(stack) == 0 || heights[i] > heights[stack[top]] {
			stack = append(stack, i)
			i++
			continue
		}
		h = heights[stack[top]]
		stack = stack[:top]
		w = i - 1 - 0 + 1
		if len(stack) > 0 {
			top = len(stack) - 1
			w = i - 1 - stack[top]
		}
		area = h * w
		if area > max {
			max = area
		}
	}
	return max
}

func largestRectangleAreaOrdinaryMethod(heights []int) int {
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

func largestRectangleArea(heights []int) int {
	return largestRectangleAreaStackMethod(heights)
}
