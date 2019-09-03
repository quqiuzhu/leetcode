package main

/*
 * @lc app=leetcode id=155 lang=golang
 *
 * [155] Min Stack
 */

// MinStack stack
type MinStack struct {
	stack []int
	min   int
}

// NewMinStack newMinStack
func NewMinStack() MinStack {
	return MinStack{}
}

// 与146题重名， 提交时取消下列注释
/** initialize your data structure here. */
// func Constructor() MinStack {
// 	return NewMinStack()
// }

// Push push
func (s *MinStack) Push(x int) {
	if len(s.stack) == 0 || x < s.min {
		s.min = x
	}
	s.stack = append(s.stack, x)
}

// Pop pop
func (s *MinStack) Pop() {
	if s.Top() == s.min {
		s.min = s.stack[0]
		for i := 0; i < len(s.stack)-1; i++ {
			if s.stack[i] < s.min {
				s.min = s.stack[i]
			}
		}
	}
	s.stack = s.stack[:len(s.stack)-1]
}

// Top top
func (s *MinStack) Top() int {
	return s.stack[len(s.stack)-1]
}

// GetMin getMin
func (s *MinStack) GetMin() int {
	return s.min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
