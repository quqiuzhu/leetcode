package main

/*
 * @lc app=leetcode id=85 lang=golang
 *
 * [85] Maximal Rectangle
 * 每一行记录前面(包含本身)有几个连续的1，然后对每一列处理
 * 每一列最大的面积就是求最大直方图面积，即84题所求
 * 行列反过来也一样，为方便最大直方图计算，反过来
 */

// maxHistArea copy from 84
func maxHistArea(heights []int) int {
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

func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	m, n := len(matrix), len(matrix[0])
	var i, j, v, max, area int
	h := make([][]int, m)
	for i = 0; i < m; i++ {
		h[i] = make([]int, n)
	}

	for i = 0; i < m; i++ {
		for j = 0; j < n; j++ {
			if matrix[i][j] == '0' {
				v = 0
			} else if i == 0 {
				v = 1
			} else {
				v = h[i-1][j] + 1
			}
			h[i][j] = v
		}
	}

	for i = 0; i < m; i++ {
		area = maxHistArea(h[i])
		if area > max {
			max = area
		}
	}
	return max
}
