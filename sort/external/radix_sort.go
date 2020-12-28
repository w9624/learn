package main

// 多关键字排序，O(d*n)
func RadixSort(array []int) {
	if len(array) <= 1 {
		return
	}

	max := array[0]
	for _, val := range array {
		if max < val {
			max = val
		}
	}

	keyLen := 0
	for max > 0 {
		max = max / 10
		keyLen++
	}

	// todo: 需要两个同样的结构，其中一个在循环中做临时数据存储
	bucket := make([][]int, 10)
	for i := 0; i < keyLen; i++ {
		_ = bucket
	}
}
