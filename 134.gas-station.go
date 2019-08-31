package main

/*
 * @lc app=leetcode id=134 lang=golang
 *
 * [134] Gas Station
 */
func canCompleteCircuit(gas []int, cost []int) int {
	var start, maxStart, i, cur, max int
	cur = gas[0] - cost[0]
	max = cur
	for i = 1; i < len(cost); i++ {
		if cur > 0 {
			cur += gas[i] - cost[i]
		} else {
			start = i
			cur = gas[i] - cost[i]
		}
		if cur >= max {
			maxStart = start
		}
		gas[i] += gas[i-1]
		cost[i] += cost[i-1]
	}

	var gasBase, costBase, gasRemain, costRemain int
	if maxStart > 0 {
		gasBase = gas[maxStart-1]
		costBase = cost[maxStart-1]
	}
	gasRemain = gas[len(gas)-1] - gasBase
	costRemain = cost[len(cost)-1] - costBase
	for i = 0; i < len(cost); i++ {
		if i >= maxStart {
			cost[i] -= costBase
			gas[i] -= gasBase
		} else {
			cost[i] += costRemain
			gas[i] += gasRemain
		}
		if gas[i]-cost[i] < 0 {
			return -1
		}
	}
	return maxStart
}
