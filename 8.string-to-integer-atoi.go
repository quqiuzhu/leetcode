package main

/*
 * @lc app=leetcode id=8 lang=golang
 *
 * [8] String to Integer (atoi)
 */

// State nolint
type State int

// nolint
const (
	InitState         State = iota
	BlankState        State = iota
	SignState         State = iota
	ProcessState      State = iota
	EndErrorState     State = iota
	EndOverfolowState State = iota
	EndNormalState    State = iota
	MaxDivideTen      int   = 0x07fffffff / 10
)

func myAtoi(str string) int {
	state := InitState
	var sign, value int

	for {
		switch state {
		case InitState:
			sign, value = 0, 0
			if len(str) > 0 {
				state = BlankState
			} else {
				state = EndErrorState
			}
		case BlankState:
			if len(str) == 0 {
				state = EndErrorState
			} else if str[0] == ' ' {
				str = str[1:]
				state = BlankState
			} else {
				state = SignState
			}
		case SignState:
			if str[0] == '-' {
				sign = -1
				str = str[1:]
			} else if str[0] == '+' {
				sign = 1
				str = str[1:]
			} else {
				sign = 1
			}
			state = ProcessState
		case ProcessState:
			if len(str) > 0 && str[0] >= '0' && str[0] <= '9' {
				newValue := value*10 + int(str[0]-'0')
				if value > MaxDivideTen {
					state = EndOverfolowState
				} else if newValue < 0 || newValue > 0x07fffffff {
					state = EndOverfolowState
				} else {
					value = newValue
					str = str[1:]
					state = ProcessState
				}
			} else {
				state = EndNormalState
			}
		case EndErrorState:
			return 0
		case EndOverfolowState:
			if sign < 0 {
				return -0x80000000
			}
			return 0x07fffffff
		case EndNormalState:
			return value * sign
		}
	}
}
