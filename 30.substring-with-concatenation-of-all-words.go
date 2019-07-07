package main

/*
 * @lc app=leetcode id=30 lang=golang
 *
 * [30] Substring with Concatenation of All Words
 */
func findSubstring(s string, words []string) []int {
	var results []int
	ls, lc := len(s), len(words)
	if lc == 0 {
		return results
	}
	lw := len(words[0])
	if lw == 0 {
		return results
	}

	wCounts := make(map[string]int)
	for _, w := range words {
		if _, ok := wCounts[w]; ok {
			wCounts[w]++
		} else {
			wCounts[w] = 1
		}
	}

	var i, j int
	var w string
	var check bool
	for i = 0; i <= ls-lw*lc; i++ {
		counts := make(map[string]int)

		check = true
		for j = 0; j < lc; j++ {
			w = s[i+j*lw : i+(j+1)*lw]
			if _, ok := wCounts[w]; !ok {
				check = false
				break
			}
			if _, ok := counts[w]; ok {
				counts[w]++
			} else {
				counts[w] = 1
			}
		}

		if check {
			for k, v := range wCounts {
				if v != counts[k] {
					check = false
				}
			}
			if check {
				results = append(results, i)
			}
		}
	}

	return results
}
