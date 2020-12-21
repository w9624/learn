/*
 * Copyright (C) Runze Wang
 * Copyright (C) ixiaochuan.cn
 */

package main

import "fmt"

// 最小的K个数
// topk问题
// 1. 冒泡O(kn) 2. 堆排O(k logn) 3. RandomSelect O(n)(快排)
func GetLeastNumbers(array []int, k int) []int {

	if len(array) == 0 {
		return nil
	} else if len(array) <= k {
		return array
	}

	// 冒泡
	//for i := 0; i < k; i++ {
	//	exchange := true
	//	for j := len(array) - i - 1; j > 0; j-- {
	//		if array[j-1] > array[j] {
	//			array[j-1], array[j] = array[j], array[j-1]
	//			exchange = false
	//		}
	//	}
	//
	//	if exchange {
	//		break
	//	}
	//}

	// 堆排 有点费事有时间写

	a := RandomSelect(array, 0, len(array)-1, k)
	fmt.Println(a, " ...")

	return array[:k]
}

func RandomSelect(array []int, low, high, k int) int {

	if low == high {
		return low
	}

	mid := partition(array, low, high)
	if mid-low >= k {
		return RandomSelect(array, low, mid-1, k)
	} else {
		return RandomSelect(array, mid+1, high, k-mid)
	}
}

func partition(array []int, low, high int) int {

	pivot := array[low]

	for low < high {
		for low < high && pivot <= array[high] {
			high--
		}

		if low < high {
			array[low], array[high] = array[high], array[low]
		}

		for low < high && pivot > array[low] {
			low++
		}

		if low < high {
			array[low], array[high] = array[high], array[low]
		}
	}
	array[low] = pivot

	return low
}

//func build(array []int) {
//	for i := (len(array) - 1) / 2; i >= 0; i-- {
//		rebuild(array, i, len(array))
//	}
//}
//
//func rebuild(array []int)  {
//
//}

func main() {
	fmt.Println(GetLeastNumbers([]int{4, 5, 1, 6, 2, 7, 3, 1, 8}, 4))
}
