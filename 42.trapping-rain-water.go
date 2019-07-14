package main

/*
 * @lc app=leetcode id=42 lang=golang
 *
 * [42] Trapping Rain Water
 */
func trap(height []int) int {
	var stack []int
	var i, j, ls, top, bottom, width, area int
	for i = 0; i < len(height); i++ {
		ls = len(stack)
		if ls == 0 || height[stack[ls-1]] >= height[i] {
			stack = append(stack, i)
			continue
		}

		bottom = height[stack[ls-1]]
		stack = stack[:ls-1]
		ls = len(stack)
		for ls > 0 {
			j = stack[ls-1]
			top = height[j]
			if top >= height[i] {
				top = height[i]
			}
			width = i - j - 1
			// fmt.Println(height[i], height[j], width, top, bottom, width*(top-bottom))
			area += width * (top - bottom)
			if height[j] > height[i] {
				break
			}

			bottom = top
			stack = stack[:ls-1]
			ls = len(stack)
		}
		stack = append(stack, i)
	}
	return area
}
