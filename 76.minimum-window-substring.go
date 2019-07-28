package main

/*
 * @lc app=leetcode id=76 lang=golang
 *
 * [76] Minimum Window Substring
 */
func minWindow(s string, t string) string {
	var i, j, v, count, window int
	var ok bool
	var result string

	tm := make(map[byte]int)
	sm := make(map[byte]int)

	for i = 0; i < len(t); i++ {
		tm[t[i]]++
	}

	for i = 0; i < len(s); i++ {
		if v, ok = tm[s[i]]; ok {
			sm[s[i]]++
			if sm[s[i]] <= v {
				count++
			}
		}
		if count < len(t) {
			continue
		}

		for ; j <= i; j++ {
			if _, ok = tm[s[j]]; !ok {
				continue
			}
			if sm[s[j]] > tm[s[j]] {
				sm[s[j]]--
				continue
			}
			break
		}
		if window == 0 || i+1-j < window {
			result = s[j : i+1]
			window = i + 1 - j
		}
		sm[s[j]]--
		count--
		j++
	}
	return result
}
