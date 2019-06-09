package main

/*
 * @lc app=leetcode id=7 lang=golang
 *
 * [7] Reverse Integer
 */

func sign(x int) int {
	s := 0
	if x > 0 {
		s = 1
	} else if x < 0 {
		s = -1
	}
	return s
}

func reverse(x int) int {
	// handle zero & -2^31
	min := -0x80000000
	if x == 0 || x == min {
		// fmt.Println("min:", min)
		return 0
	}

	// sign
	s := sign(x)
	max := 0x07fffffff / 10
	// fmt.Println("max:", max)
	reversed := 0
	x = x * s
	for x != 0 {
		if reversed > max {
			// fmt.Println("reversed:", reversed, "max:", max)
			return 0
		}
		remainder := x % 10
		tmp := reversed*10 + remainder
		// fmt.Println("tmp:", tmp)
		if tmp < 0 {
			return 0
		}
		reversed = tmp
		x = x / 10
	}
	return reversed * s
}
