package main

import "log"

// 跳台阶
// 斐波拉契数列，不过是以1开头
func JumpFloor(target int) int {

	if target == 0 {
		return 0
	} else if target == 1 {
		return 1
	} else if target == 2 {
		return 2
	}

	//result := 0
	//a, b := 1, 2
	//for i := 3; i <= target; i++ {
	//	result = a + b
	//	a = b
	//	b = result
	//}
	//
	//return result

	return JumpFloor(target-1) + JumpFloor(target-2)
}

func main() {
	log.Print(JumpFloor(4))
}
