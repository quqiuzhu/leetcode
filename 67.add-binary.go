package main

/*
 * @lc app=leetcode id=67 lang=golang
 *
 * [67] Add Binary
 */
func addBinary(a string, b string) string {
	if len(a) < len(b) {
		a, b = b, a
	}

	var x, y []byte
	x, y = []byte(a), []byte(b)
	var r, v, m, n byte
	la, lb := len(a), len(b)
	for i := 0; i < la; i++ {
		m = x[la-1-i] - '0'
		n = 0
		if i < lb {
			n = y[lb-1-i] - '0'
		}
		v = m + n + r
		r = v / 2
		x[la-1-i] = v%2 + '0'
	}

	if r == 0 {
		return string(x)
	}

	result := make([]byte, la+1)
	result[0] = '1'
	copy(result[1:], x)

	return string(result)
}
