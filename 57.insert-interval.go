package main

/*
 * @lc app=leetcode id=57 lang=golang
 *
 * [57] Insert Interval
 */

func insertPosition(intervals [][]int, x int) int {
	var mid, begin, end, position int
	begin, end = 0, len(intervals)-1
	for begin <= end {
		mid = (begin + end) / 2
		if x == intervals[mid][0] {
			position = mid + 1
			break
		}
		if x > intervals[mid][0] {
			position = mid + 1
			begin = mid + 1
		} else {
			position = mid
			end = mid - 1
		}
	}
	return position
}

func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}

	x := insertPosition(intervals, newInterval[0])
	y := insertPosition(intervals, newInterval[1])

	var r [][]int
	var first, second int
	first = newInterval[0]
	second = newInterval[1]
	if x != 0 && intervals[x-1][1] >= first {
		first = intervals[x-1][0]
		x--
	}
	if y != 0 && intervals[y-1][1] > second {
		second = intervals[y-1][1]
	}

	for i := 0; i < x; i++ {
		r = append(r, intervals[i])
	}
	r = append(r, []int{first, second})
	for i := y; i < len(intervals); i++ {
		r = append(r, intervals[i])
	}

	return r
}
