package main

import "log"

// 调整数组顺序使奇数位于偶数前
// 排序 -> 插入排序 -> 排序稳定性
func reOrderArray(array []int) []int {

	if len(array) <= 1 {
		return array
	}

	for i := 1; i < len(array); i++ {
		if array[i]%2 == 1 {
			cur := array[i]
			j := i
			for j > 0 && array[j-1]%2 == 0 {
				array[j] = array[j-1]
				j--
			}
			array[j] = cur
		}
	}

	return array
}

func main() {
	log.Println(reOrderArray([]int{3, 2, 5, 1, 3, 4, 7, 9, 8, 2}))
}
