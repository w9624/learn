package main

import "log"

// 孩子们的游戏
// 约瑟夫环问题
// 1. 数组/链表 -> 环
// 2. 公式 f(n,m) = (f(n−1,m) + m) % n
func LastRemaining(n int, m int) int {
	if n <= 1 {
		return -1
	}

	arr := make([]int, n)
	free := n
	index := m - 1
	for free > 1 {
		for arr[index%n] == -1 {
			index++
			index %= n
		}
		log.Println(index)
		arr[index%n] = -1
		index = (index + m) % n
		free--
		log.Println(arr)
	}

	log.Println(arr)
	for i, val := range arr {
		if val == 0 {
			return i
		}
	}
	return -1

	// 公式
	//index = 0
	//for i := 2; i <=n; i++ {
	//	index = (index + m) % i
	//}
	//
	//return index
}

func main() {
	log.Println(LastRemaining(5, 3))
}
