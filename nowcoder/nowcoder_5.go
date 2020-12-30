package main

// 用两个栈实现队列
// 栈和队列
import (
	"github.com/golang-collections/collections/stack"
)

var s1 stack.Stack
var s2 stack.Stack

func Push(node int) {
	for s1.Len() > 0 {
		s2.Push(s1.Pop())
	}
	s1.Push(node)
	for s2.Len() > 0 {
		s1.Push(s2.Pop())
	}
}

func Pop() int {
	if s1.Len() > 0 {
		return s1.Pop().(int)
	}

	return -1
}
