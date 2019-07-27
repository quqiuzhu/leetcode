package main

/*
 * @lc app=leetcode id=66 lang=golang
 *
 * [66] Plus One
 */
func plusOne(digits []int) []int {
	r, v := 1, 0
	for i := len(digits) - 1; i >= 0; i-- {
		v = digits[i] + r
		r = v / 10
		digits[i] = v % 10
		if r == 0 {
			break
		}
	}

	d := digits
	if r != 0 {
		d = make([]int, len(digits)+1)
		d[0] = r
		copy(d[1:], digits)
	}
	return d
}
