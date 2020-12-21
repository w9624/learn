package main

import "fmt"

// 数组中出现次数超过一半的数字
// 别搞哪些花里胡哨的, 一个map解决问题....
func MoreThanHalfNum(array []int) int {

	if len(array) == 0 {
		return -1
	}

	key, maxCount := array[0], 0
	m := make(map[int]int, 0)
	for _, i := range array {
		m[i]++
		if m[i] > maxCount {
			maxCount = m[i]
			key = i
		}
	}

	return key
}

func main() {
	fmt.Println(MoreThanHalfNum([]int{1, 2, 3, 2, 2, 2, 5, 4, 2}))
}
