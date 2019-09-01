package main

/*
 * @lc app=leetcode id=149 lang=golang
 *
 * [149] Max Points on a Line
 */

// y = ax + b
// y = (au/ab)x + bu/bb
// u: upper b: bottom
// when ab is zero, record x
type line struct {
	au, ab, bu, bb int
}

type pointNum struct {
	num, lastIndex int
}

func simplifyDivisor(m, n int) (int, int) {
	if m == 0 && n == 0 {
		return m, n
	}
	if m == 0 {
		return m, 1
	}
	if n == 0 {
		return 1, n
	}

	var a, b int
	a = n
	b = m
	for b != 0 {
		a, b = b, a%b
	}
	m /= a
	n /= a
	return m, n
}

func maxPoints(points [][]int) int {
	if len(points) == 0 {
		return 0
	}

	var i, j, l, au, ab, bu, bb int
	lineMap := make(map[line]*pointNum)
	l = len(points)

	var pn *pointNum
	var ln line
	same := true
	for i = 0; i < l; i++ {
		for j = 0; j < l; j++ {
			if points[i][0] == points[j][0] && points[i][1] == points[j][1] {
				continue
			}
			same = false
			ab, au = simplifyDivisor(points[i][0]-points[j][0], points[i][1]-points[j][1])
			if ab == 0 {
				bb = 0
				bu = points[i][0]
			} else {
				bb = ab
				bu = ab*points[i][1] - au*points[i][0]
				bb, bu = simplifyDivisor(bb, bu)
			}

			ln = line{au: au, ab: ab, bu: bu, bb: bb}
			pn = lineMap[ln]
			if pn == nil {
				lineMap[ln] = &pointNum{
					num:       1,
					lastIndex: i,
				}
				continue
			}
			if pn.lastIndex != i {
				pn.num++
				pn.lastIndex = i
			}
		}
	}

	if same {
		return len(points)
	}

	max := 0
	for ln := range lineMap {
		if lineMap[ln].num > max {
			max = lineMap[ln].num
		}
	}
	return max
}
