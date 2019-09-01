package main

import (
	"strconv"
)

/*
 * @lc app=leetcode id=150 lang=golang
 *
 * [150] Evaluate Reverse Polish Notation
 */
func evalRPN(tokens []string) int {
	var stack []int
	var x, y int
	for i := 0; i < len(tokens); i++ {
		if tokens[i] == "+" || tokens[i] == "-" ||
			tokens[i] == "*" || tokens[i] == "/" {
			x, y = stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			switch tokens[i] {
			case "+":
				stack = append(stack, x+y)
			case "-":
				stack = append(stack, x-y)
			case "*":
				stack = append(stack, x*y)
			case "/":
				stack = append(stack, x/y)
			}
		} else {
			x, _ = strconv.Atoi(tokens[i])
			stack = append(stack, x)
		}
	}
	return stack[len(stack)-1]
}
