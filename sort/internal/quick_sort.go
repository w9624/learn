/*
 * Copyright (C) Runze Wang
 * Copyright (C) ixiaochuan.cn
 */

package main

import (
	"log"
	"sort"
)

func QuickSort(array []int) {
	quickSort(array, 0, len(array)-1)
}

func quickSort(array []int, low, high int) {
	for low < high {
		mid := partition(array, low, high)
		quickSort(array, low, mid-1)
		low = mid + 1
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
	array := []int{3, 2, 2, 9, 4, 6}
	QuickSort(array)
	log.Printf("sorted: %v \narray: %v", sort.IntsAreSorted(array), array)
}
