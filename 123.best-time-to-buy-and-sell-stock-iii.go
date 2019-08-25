package main

/*
 * @lc app=leetcode id=123 lang=golang
 *
 * [123] Best Time to Buy and Sell Stock III
 */
func maxProfitIII(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	profitsI := make([]int, len(prices))
	profitsII := make([]int, len(prices))
	var i, profit, min, max, v int

	min = prices[0]
	for i = 1; i < len(prices); i++ {
		v = prices[i] - min
		if v > profit {
			profit = v
		}
		if prices[i] < min {
			min = prices[i]
		}
		profitsI[i] = profit
	}

	profit = 0
	max = prices[len(prices)-1]
	for i = len(prices) - 2; i >= 0; i-- {
		v = max - prices[i]
		if v > profit {
			profit = v
		}
		if prices[i] > max {
			max = prices[i]
		}
		profitsII[i] = profit
	}

	profit = profitsI[len(prices)-1]
	for i = 0; i < len(prices)-1; i++ {
		v = profitsI[i] + profitsII[i+1]
		if v > profit {
			profit = v
		}
	}
	return profit
}

// 与121重名， 提交时取消下列注释
// func maxProfit(prices []int) int {
// 	return maxProfitIII(prices)
// }
