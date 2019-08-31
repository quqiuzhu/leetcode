package main

/*
 * @lc app=leetcode id=135 lang=golang
 *
 * [135] Candy
 */
func candy(ratings []int) int {
	l := len(ratings)
	if l <= 1 {
		return l
	}
	var i, total, cur, peak, peakVal int
	for i = 1; i < len(ratings); i++ {
		if ratings[i] == ratings[i-1] {
			break
		}
		if ratings[i] > ratings[i-1] {
			cur++
			total += cur
			peak = i
			peakVal = cur
		} else {
			cur = 0
			total += i - peak - 1
			if i-peak-1 >= peakVal {
				total++
			}
		}
	}
	total += i
	if i < len(ratings) {
		total += candy(ratings[i:])
	}

	return total
}
