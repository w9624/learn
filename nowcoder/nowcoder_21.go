package main

import (
	"fmt"
	"github.com/golang-collections/collections/stack"
)

// 包含min函数的栈
// 这道题的重点在于，栈中的数是后进先出（即出入顺序相反）
// md，用链表解了保存了所有值，每次push需要O(n)...当时没想到啊...
// 还有个问题是go中没有stack，如果要写的还只能先实现stack...
type StackMinNum struct {
	S    *stack.Stack
	SMax *stack.Stack
}

func (s *StackMinNum) Push(val int) {
	if s.SMax.Len() == 0 || val >= s.SMax.Peek().(int) {
		s.SMax.Push(val)
	}
	s.S.Push(val)
}

func (s *StackMinNum) Pop() int {
	if s.S.Len() == 0 {
		return 0
	}

	val := s.S.Pop()
	if val == s.SMax.Peek() {
		s.SMax.Pop()
	}
	return val.(int)
}

func (s *StackMinNum) Min() int {
	if s.SMax.Len() == 0 {
		return 0
	}

	return s.SMax.Peek().(int)
}

func main() {
	s := StackMinNum{
		S:    stack.New(),
		SMax: stack.New(),
	}

	array := []int{2, 5, 3, 1, 6, 5, 10, 7, 9, 2}
	for _, val := range array {
		s.Push(val)
	}

	for range array {
		fmt.Println(s.Min())
		s.Pop()
	}
}
