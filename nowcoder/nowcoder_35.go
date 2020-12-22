/*
 * Copyright (C) Runze Wang
 * Copyright (C) ixiaochuan.cn
 */

package main

import "fmt"

// 数组中的逆序对
// 归并排序，时间复杂度O(n), 空间复杂度O(n), 利用partition归并的时候获得逆序对数
// 递归中传值挺麻烦的，简化的方式定义全局变量累增
func InversePairs(array []int) int {
	return sort(array, 0, len(array)-1)
}

func sort(array []int, low, high int) int {
	if low >= high {
		return 0
	}

	var count int
	mid := (low + high) / 2
	count += sort(array, low, mid)
	count += sort(array, mid+1, high)
	if array[mid] <= array[mid+1] {
		return count
	}
	count += merge(array, low, mid, high)
	return count
}

func merge(array []int, low, mid, high int) int {
	tmp := make([]int, high-low+1)

	var count int
	i, j, k := low, mid+1, 0
	for i <= mid && j <= high {
		if array[i] < array[j] {
			tmp[k] = array[i]
			i++
		} else {
			count = (count + mid + 1 - i) % 1000000007
			tmp[k] = array[j]
			j++
		}
		k++
	}

	for i <= mid {
		tmp[k] = array[i]
		i++
		k++
	}

	for j <= high {
		tmp[k] = array[j]
		j++
		k++
	}

	for i := 0; i < len(tmp); i++ {
		array[low+i] = tmp[i]
	}

	return count
}

func main() {
	fmt.Println(InversePairs([]int{1, 2, 3, 4, 5, 6, 7, 0}))
}
