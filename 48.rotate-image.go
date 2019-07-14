package main

/*
 * @lc app=leetcode id=48 lang=golang
 *
 * [48] Rotate Image
 * 从外圈到内圈依次旋转
 */

func idx(c, index, l int) (int, int) {
	var i, j, k int
	width := l - 2*c - 1
	k = index / width
	if k == 0 {
		i = c
		j = c + index%width
	} else if k == 1 {
		j = l - 1 - c
		i = c + index%width
	} else if k == 2 {
		i = l - 1 - c
		j = c + width - (index % width)
	} else {
		j = c
		i = c + width - (index % width)
	}
	return i, j
}

func rotate(matrix [][]int) {
	l := len(matrix)
	var c, index, width, cl, i, j, k, h, tmp, f int
	for c = 0; c < l/2; c++ {
		width = l - 2*c - 1
		cl = width * 4
		for f = 0; f < width; f++ {
			i, j = idx(c, cl-1, l)
			tmp = matrix[i][j]
			for index = cl - 1; index > 0; index-- {
				i, j = idx(c, index-1, l)
				// fmt.Println(matrix[i][j])
				k, h = idx(c, index, l)
				matrix[k][h] = matrix[i][j]
			}
			i, j = idx(c, 0, l)
			matrix[i][j] = tmp
		}
	}
}
