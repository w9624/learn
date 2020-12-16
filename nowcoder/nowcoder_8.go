package main

import "log"

// 跳台阶
func JumpFloor(target int) int {

	if target == 0 {
		return 0
	} else if target == 1 {
		return 1
	} else if target == 2 {
		return 2
	}

	//a, b, c := 1, 2, 0
	//for i := 3; i <= target; i++ {
	//	c = a + b
	//	a = b
	//	b = c
	//}
	//
	//return c

	return JumpFloor(target-1) + JumpFloor(target-2)
}

func main() {
	log.Print(JumpFloor(4))
}
