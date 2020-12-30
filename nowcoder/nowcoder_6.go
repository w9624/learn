package main

import "log"

// 旋转数组最小数字
// 二分查找 -> 边界划分
/*
1. mid > high, low = mid + 1, 右侧
2. mid = high, high = high - 1, 无法判断最小值在mid左侧还是右侧
3. mid < high, high = mid，可能为当前值因此high=mid
*/
func MinNumberInRotateArray(array []int) int {
	if len(array) == 0 {
		return 0
	}

	low, high := 0, len(array)-1

	for low < high {
		mid := (low + high) / 2
		if array[mid] > array[high] {
			low = mid + 1
		} else if array[mid] == array[high] {
			high = high - 1
		} else {
			high = mid
		}
		// 2,3,4,5,1,2
		// log.Printf("low: %v, high: %v", low, high)
	}

	return array[low]
}

func main() {
	cases := []struct {
		array    []int
		expected int
	}{
		{
			array:    []int{3, 4, 5, 1, 2},
			expected: 1,
		},
		{
			array:    []int{2, 3, 4, 5, 1, 2},
			expected: 1,
		},
		{
			array:    []int{4, 5, 1, 2, 3, 3, 3, 3, 3},
			expected: 1,
		},
		{
			array:    []int{1, 0, 1, 1, 1, 1},
			expected: 0,
		},
	}

	for _, c := range cases {
		if min := MinNumberInRotateArray(c.array); c.expected != min {
			log.Fatalf("expected: %v, mid: %v", c.expected, min)
		} else {
			log.Println(min)
		}

	}
}
