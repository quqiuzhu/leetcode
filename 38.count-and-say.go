package main

/*
 * @lc app=leetcode id=38 lang=golang
 *
 * [38] Count and Say
 */

func countAndSayString(n int, s string) string {
	if n <= 1 {
		return s
	}

	var bs []byte
	b := byte('0')
	count := 0
	for i := 0; i < len(s); i++ {
		if b != s[i] {
			if count != 0 {
				bs = append(bs, byte('0'+count))
				bs = append(bs, b)
			}

			count = 1
			b = s[i]
		} else {
			count++
		}
	}
	if count != 0 {
		bs = append(bs, byte('0'+count))
		bs = append(bs, b)
	}

	return countAndSayString(n-1, string(bs))
}

func countAndSay(n int) string {
	return countAndSayString(n, "1")
}
