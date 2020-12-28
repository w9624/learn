package main

import (
	"log"
)

// 常见解法
// 1. 快排 T(nlogn)
// 2. 堆排 T(klogn)
// 3. RS T(n)
// 海量数据解法
// 1. 每个机器维持小/大根堆，所有机器的堆顶元素再组成小/大根堆
func Topk(array []int, k int) {
	if len(array) == 0 || k == 0 || len(array) < k {
		return
	}

	// quick sort

	// heap sort

	// rs
	randomSelect(array, 0, len(array)-1, k)

	return
}

func randomSelect(array []int, low, high, k int) int {
	if low == high {
		return array[low]
	}

	mid := partition(array, low, high)
	index := mid - low + 1
	if index == k {
		return array[low]
	} else if index > k {
		return randomSelect(array, low, mid-1, k)
	} else {
		return randomSelect(array, mid+1, high, k-mid)
	}
}

func partition(array []int, low, high int) int {
	pivot := array[low]

	for low < high {
		for low < high && pivot <= array[high] {
			high--
		}

		if low < high {
			array[low] = array[high]
		}

		for low < high && pivot > array[low] {
			low++
		}

		if low < high {
			array[high] = array[low]
		}
	}

	array[low] = pivot
	return low
}

func main() {
	array := []int{3, 2, 2, 3, 9, 4, 6, 5, 1, 4}
	Topk(array, 4)
	log.Println(array[:3])
}
