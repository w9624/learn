package main

import (
	"log"

	"github.com/badgerodon/collections/stack"
)

// 栈的压入、弹出序列
// 栈 -> 判处出栈序列合法性
// 1. 顺序入栈，
// 2. 栈不为空，比较栈顶元素和出栈数组当前元素，相等则出栈且数组元素后移
// 3. 判断栈是否为空
func IsPopOrder(pushV []int, popV []int) bool {
	s := stack.Stack{}

	j := 0
	for _, val := range pushV {
		s.Push(val)
		for s.Len() > 0 && s.Peek() == popV[j] {
			s.Pop()
			j++
		}
	}

	return s.Len() == 0
}

func main() {
	log.Printf("%v", IsPopOrder([]int{1, 2, 3, 4, 5}, []int{4, 5, 3, 2, 1}))
}
