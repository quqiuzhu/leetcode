package main

/*
 * @lc app=leetcode id=71 lang=golang
 *
 * [71] Simplify Path
 */
func simplifyPath(path string) string {
	var stack []string
	var name string
	p := 0
	for i := 0; i < len(path); i++ {
		if path[i] != '/' && i != len(path)-1 {
			continue
		}

		name = ""
		if path[i] == '/' {
			if i-p > 1 {
				name = path[p+1 : i]
			}
		} else if i == len(path)-1 {
			name = path[p+1 : len(path)]
		}
		p = i

		if name == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else if name != "." && len(name) > 0 {
			stack = append(stack, name)
		}
	}
	var simplified []byte
	for i := 0; i < len(stack); i++ {
		simplified = append(simplified, '/')
		for j := 0; j < len(stack[i]); j++ {
			simplified = append(simplified, stack[i][j])
		}
	}
	if simplified == nil {
		simplified = append(simplified, '/')
	}
	return string(simplified)
}
