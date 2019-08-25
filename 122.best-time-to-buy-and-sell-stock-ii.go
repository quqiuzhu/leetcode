package main

/*
 * @lc app=leetcode id=122 lang=golang
 *
 * [122] Best Time to Buy and Sell Stock II
 */
func maxProfitII(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	var i, profit, v int
	for i = 1; i < len(prices); i++ {
		v = prices[i] - prices[i-1]
		if v > 0 {
			profit += v
		}
	}
	return profit
}

// 与121题重名， 提交时取消下列注释
// func maxProfit(prices []int) int {
// 	return maxProfitII(prices)
// }
