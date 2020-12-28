package main

import (
	"fmt"
	"log"
	"sort"
)

func CountSort(array []int) {
	if len(array) <= 1 {
		return
	}

	min, max := array[0], array[0]
	for _, i := range array {
		if min > i {
			min = i
		}
		if max < i {
			max = i
		}
	}

	count := max - min + 1
	buckets := make([]int, count)
	for _, val := range array {
		bIndex := val - min
		buckets[bIndex]++
	}

	fmt.Println(buckets)

	k := 0
	for index, val := range buckets {
		for val > 0 {
			array[k] = index + min
			val--
			k++
		}
	}
}

func main() {
	array := []int{3, 2, 9, 4, 6}
	CountSort(array)
	log.Printf("sorted: %v \narray: %v", sort.IntsAreSorted(array), array)
}
