package main

import (
	"fmt"
	"strconv"
)

// 数组组成的最小数字
// 关键点在于 想到比较两个拼接后字符串的大小，并不是拼接前的大小
func PrintMinNumber(array []int) string {
	if len(array) == 0 {
		return ""
	} else if len(array) == 1 {
		return strconv.Itoa(array[0])
	}

	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			s1 := strconv.Itoa(array[i]) + strconv.Itoa(array[j])
			s2 := strconv.Itoa(array[j]) + strconv.Itoa(array[i])
			if s1 > s2 {
				array[i], array[j] = array[j], array[i]
			}
		}
	}

	var result string
	for _, i := range array {
		result = result + strconv.Itoa(i)
	}

	return result
}

func main() {
	fmt.Println(PrintMinNumber([]int{3, 32, 321}))
}
