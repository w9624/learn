/*
 * Copyright (C) Runze Wang
 * Copyright (C) ixiaochuan.cn
 */

package main

import "fmt"

// 数组中只出现一次的数字
// 1. 常规逻辑使用map统计出现次数
// 2. 使用逻辑运算，c=a^b -> a=c^b, b=a^c，同时需要划分数组将a和b划分在不同的数组(因a、b不同因此c定有一位bit位为1)，通过c恢复
// tmd，map不香吗，非要搞这么多花里胡哨的，不就省下O(m)的空间吗...
func FindNumsAppearOnce(array []int) (num1, num2 []int) {
	if len(array) == 0 {
		return
	} else if len(array) == 1 {
		num1 = append(num1, array[0])
	} else if len(array) == 2 {
		num1 = append(num1, array[0])
		num2 = append(num2, array[1])
		return
	}

	tmp := 0
	for _, val := range array {
		tmp ^= val
	}
	num1 = append(num1, tmp)
	num2 = append(num2, tmp)

	var index uint
	for (tmp&1 == 0) && index < 32 {
		tmp >>= 1
		index++
	}

	isSameArr := func(val int, index uint) bool {
		val >>= index
		return val&1 == 1
	}
	for _, val := range array {
		if isSameArr(val, index) {
			num1[0] ^= val
		} else {
			num2[0] ^= val
		}
	}

	return
}

func main() {
	fmt.Println(FindNumsAppearOnce([]int{3, 4, 3, 3, 3, 3, 2, 3}))
}
