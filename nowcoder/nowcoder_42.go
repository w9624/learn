package main

import (
	"log"
	"math"
)

// 和为S的两个数字
// 数组 -> 有序数组 -> 双指针法
func FindNumbersWithSum(array []int, target int) []int {
	if len(array) < 2 {
		return nil
	}

	mutilply := math.MaxInt32
	result := make([]int, 2)
	i, j := 0, len(array)-1
	for i < j {
		sum := array[i] + array[j]
		if sum == target && array[i]*array[j] < mutilply {
			result[0], result[1] = array[i], array[j]
			mutilply = array[i] * array[j]
			i++
			j--
		} else if sum > target {
			j--
		} else {
			i++
		}
	}

	return result
}

func main() {
	log.Println(FindNumbersWithSum([]int{1, 2, 4, 7, 8, 12, 14, 15}, 16))
}
