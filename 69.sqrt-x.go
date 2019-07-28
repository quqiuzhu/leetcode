package main

/*
 * @lc app=leetcode id=69 lang=golang
 *
 * [69] Sqrt(x)
 */
func mySqrt(x int) int {
	var begin, end, mid, v, result int
	end = x
	for begin <= end {
		mid = (begin + end) / 2
		v = mid * mid
		if v == x {
			result = mid
			break
		}
		if v > x {
			end = mid - 1
		} else {
			result = mid
			begin = mid + 1
		}
	}
	return result
}
