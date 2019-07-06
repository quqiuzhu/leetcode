package main

/*
 * @lc app=leetcode id=20 lang=golang
 *
 * [20] Valid Parentheses
 */
func isValid(s string) bool {
	stack := make([]byte, 0, 5)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack = append(stack, s[i])
		} else if s[i] == ')' || s[i] == ']' || s[i] == '}' {
			var pop byte
			if len(stack) != 0 {
				pop = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
			}
			if s[i] == ')' && pop != '(' || s[i] == ']' && pop != '[' ||
				s[i] == '}' && pop != '{' {
				return false
			}
		}
	}
	return len(stack) == 0
}
