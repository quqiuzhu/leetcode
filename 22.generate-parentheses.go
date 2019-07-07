package main

/*
 * @lc app=leetcode id=22 lang=golang
 *
 * [22] Generate Parentheses
 */

func findNthC(p []byte, n int) int {
	j, r := 0, -1
	for i := 0; i < len(p); i++ {
		if p[i] == '(' {
			j++
		}
		if j == n {
			r = i + 1
			break
		}
	}
	return r
}

func generateParenthesis(n int) []string {
	if n <= 0 {
		return nil
	}

	ps := [][]byte{[]byte{'(', ')'}}
	var newPs [][]byte
	var i, j int
	for i = 1; i < n; i++ {
		for _, p := range ps {
			j = findNthC(p, i)
			for ; j <= 2*i; j++ {
				newP := make([]byte, 2*(i+1))
				copy(newP, p[:j])
				copy(newP[j:], []byte{'(', ')'})
				copy(newP[j+2:], p[j:])
				newPs = append(newPs, newP)
			}
		}
		ps = newPs
		newPs = nil
	}

	var rs []string
	for _, p := range ps {
		rs = append(rs, string(p))
	}
	return rs
}
