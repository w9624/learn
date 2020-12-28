package main

import (
	"log"
	"sort"
)

func BucketSort(array []int) {

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

	count := (max-min)/len(array) + 1
	buckets := make([][]int, count)
	// 可使用提前创建数组，避免append内存复制
	//indexs := make([]int, max-min+1)

	for _, val := range array {
		bIndex := (val - min) / len(array)
		buckets[bIndex] = append(buckets[bIndex], val)
		//buckets[bIndex][indexs[bIndex]] = val
		//indexs[bIndex]++
	}

	for _, bucket := range buckets {
		sort.Ints(bucket)
	}

	k := 0
	for _, bucket := range buckets {
		//for j := 0; j < indexs[i]; j++ {
		//	array[k] = bucket[j]
		//}
		for _, val := range bucket {
			array[k] = val
			k++
		}
	}
}

func main() {
	array := []int{3, 2, 9, 4, 6}
	BucketSort(array)
	log.Printf("sorted: %v \narray: %v", sort.IntsAreSorted(array), array)
}
