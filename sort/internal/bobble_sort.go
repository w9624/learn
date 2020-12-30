package main

import (
	"log"
	"sort"
)

// 稳定交换类排序 n^2
func BobbleSort(array []int) {
	if len(array) <= 1 {
		return
	}

	for i := 0; i < len(array)-1; i++ {
		exchange := true
		// j不对已经有序的做处理
		for j := 0; j < len(array)-1-i; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
				exchange = false
			}
		}

		if exchange {
			break
		}
	}
}

func main() {
	array := []int{3, 2, 2, 3, 9, 4, 6, 5, 1, 4}
	BobbleSort(array)
	log.Printf("sorted: %v \narray: %v", sort.IntsAreSorted(array), array)
}
