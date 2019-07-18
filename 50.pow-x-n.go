package main

/*
 * @lc app=leetcode id=50 lang=golang
 *
 * [50] Pow(x, n)
 */
func myPow(x float64, n int) float64 {
	if n < 0 {
		x = 1 / x
		n = -n
	}
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}

	m := n / 2
	powers := []float64{1.0}
	p := x
	c := 1
	i := 1
	for {
		if c >= m {
			break
		}
		p *= p
		c *= 2
		i++
	}

	power := float64(1)
	for n > 0 {
		for n >= c {
			n -= c
			power *= powers[i]
		}
		c /= 2
		i--
	}
	return power
}
