package main

import (
	"fmt"
	"sort"
)

/*
 * @lc app=leetcode id=49 lang=golang
 *
 * [49] Group Anagrams
 */

func alphaOrder(s string) string {
	b := []byte(s)
	sort.Slice(b, func(i int, j int) bool { return b[i] < b[j] })
	return string(b)
}

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]int)
	for i := range strs {
		a := alphaOrder(strs[i])
		m[a] = append(m[a], i)
	}

	fmt.Println(m)
	var rs [][]string
	for _, idxs := range m {
		var r []string
		for _, i := range idxs {
			r = append(r, strs[i])
		}
		rs = append(rs, r)
	}
	return rs
}
