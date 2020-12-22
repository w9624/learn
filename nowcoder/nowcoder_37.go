package main

import "fmt"

// 数组在升序数组中出现的次数
// 1.查找类问题 2.有序数组查找->二分查找
// 3.二分查找一个数的index，当需要知道一个数组出现次数，仅需要统计该数字两侧数字出现的下标（0-1范围内）进行做差
func GetNumberOfK(data []int, k int) int {
	if len(data) == 0 {
		return 0
	}

	search := func(data []int, k float64) int {
		low, high := 0, len(data)-1
		for low <= high { // 二分查找这里为=，如最终剩余两个数1，2，如果不是=则
			mid := (low + high) / 2
			if float64(data[mid]) <= k {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}

		return low
	}

	return search(data, float64(k)+0.5) - search(data, float64(k)-0.5)
}

func main() {
	fmt.Println(GetNumberOfK([]int{1, 2, 3, 3, 3, 3, 3, 4, 5}, 3))
}
