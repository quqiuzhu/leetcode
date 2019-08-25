package main

/*
 * @lc app=leetcode id=121 lang=golang
 *
 * [121] Best Time to Buy and Sell Stock
 */
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	var i, profit, min, v int
	min = prices[0]
	for i = 1; i < len(prices); i++ {
		v = prices[i] - min
		if v > profit {
			profit = v
		}
		if prices[i] < min {
			min = prices[i]
		}
	}
	return profit
}
