package main

import "fmt"

/*
 * @lc app=leetcode id=29 lang=golang
 *
 * [29] Divide Two Integers
 */

func abs(x int) uint32 {
	var ux uint32
	if x < 0 {
		ux = uint32(-x)
	} else {
		ux = uint32(x)
	}
	return ux
}

func divide(dividend int, divisor int) int {
	positive := dividend < 0 && divisor < 0 || dividend > 0 && divisor > 0
	d := abs(dividend)
	s := abs(divisor)

	fmt.Println(d, s)
	var x, i, j, r uint32
	x = 0x80000000
	for i = 0; i < 32; i++ {
		if s&x > 0 {
			break
		}
		s = s << 1
	}

	for j = i + 1; j > 0; j-- {
		if d >= s {
			r += x
			d -= s
		}
		x = x >> 1
		s = s >> 1
	}

	r = r >> (31 - i)
	fmt.Println(d, s, i, r)

	if r > 0x07fffffff && positive {
		return 0x07fffffff
	}

	if r > 0x80000000 && !positive {
		return -0x80000000
	}

	ri := int(r)
	if positive {
		return ri
	}

	return -ri
}
