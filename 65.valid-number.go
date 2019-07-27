package main

import "strings"

/*
 * @lc app=leetcode id=65 lang=golang
 *
 * [65] Valid Number
 * 傻逼题目，定义不清， 通过率低不是因为难
 */

type state int

const (
	signState state = iota
	intergerState
	dotStateX
	dotState
	floatState
	eState
	eIntState
	errorState
)

func isNumber(s string) bool {
	s = strings.TrimSpace(s)
	status := signState
	for len(s) != 0 {
		switch status {
		case signState:
			if s[0] == '+' || s[0] == '-' {
				s = s[1:]
			}
			if len(s) > 0 && '0' <= s[0] && s[0] <= '9' {
				status = intergerState
			} else if len(s) > 0 && s[0] == '.' {
				s = s[1:]
				status = dotStateX
			} else {
				status = errorState
			}
		case intergerState:
			if '0' <= s[0] && s[0] <= '9' {
				s = s[1:]
			} else if s[0] == '.' {
				s = s[1:]
				status = dotState
			} else if s[0] == 'e' {
				s = s[1:]
				status = eState
			} else {
				status = errorState
			}

		case dotStateX:
			if '0' <= s[0] && s[0] <= '9' {
				status = floatState
			} else {
				status = errorState
			}
		case dotState:
			if '0' <= s[0] && s[0] <= '9' {
				status = floatState
			} else if s[0] == 'e' {
				s = s[1:]
				status = eState
			} else {
				status = errorState
			}
		case floatState:
			if '0' <= s[0] && s[0] <= '9' {
				s = s[1:]
			} else if s[0] == 'e' {
				s = s[1:]
				status = eState
			} else {
				status = errorState
			}
		case eState:
			if s[0] == '+' || s[0] == '-' {
				s = s[1:]
			}
			if len(s) > 0 && '0' <= s[0] && s[0] <= '9' {
				status = eIntState
			} else {
				status = errorState
			}
		case eIntState:
			if '0' <= s[0] && s[0] <= '9' {
				s = s[1:]
			} else {
				status = errorState
			}
		case errorState:
			return false
		}
	}
	return status == intergerState ||
		status == floatState ||
		status == dotState ||
		status == eIntState
}
