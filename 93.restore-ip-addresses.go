package main

import "strings"

/*
 * @lc app=leetcode id=93 lang=golang
 *
 * [93] Restore IP Addresses
 */

func validPart(s string) bool {
	if len(s) == 0 || len(s) > 3 {
		return false
	}
	if len(s) > 1 && s[0] == '0' {
		return false
	}
	if len(s) == 3 && s > "255" {
		return false
	}
	return true
}

func restoreIpAddresses(s string) []string {
	var i, j, k, l int
	var r []string
	l = len(s)
	for i = 1; i <= 3; i++ {
		if i >= l || !validPart(s[:i]) {
			continue
		}
		for j = 1; j <= 3; j++ {
			if i+j >= l || !validPart(s[i:i+j]) {
				continue
			}
			for k = 1; k <= 3; k++ {
				if i+j+k >= l || !validPart(s[i+j:i+j+k]) {
					continue
				}
				if validPart(s[i+j+k:]) {
					var builder strings.Builder
					builder.WriteString(s[:i])
					builder.WriteByte('.')
					builder.WriteString(s[i : i+j])
					builder.WriteByte('.')
					builder.WriteString(s[i+j : i+j+k])
					builder.WriteByte('.')
					builder.WriteString(s[i+j+k:])
					r = append(r, builder.String())
				}
			}
		}
	}
	return r
}
