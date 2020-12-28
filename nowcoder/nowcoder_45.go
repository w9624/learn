package main

import (
	"log"
	"math"
)

// 扑克牌的顺子
// 技巧型问题：顺子最大和最小值做差+1<=5，同时顺子需要判读重复
// 1. 排序，除开然后做差判断<=4
// 2. 利用 顺子最大和最小值做差+1<=5的性质
func IsContinuous(numbers []int) bool {
	if len(numbers) != 5 {
		return false
	}

	m := make(map[int]int)
	min, max := math.MaxUint32, 0
	for _, val := range numbers {
		if val == 0 {
			continue
		}
		if _, ok := m[val]; ok {
			return false
		}
		if val < min {
			min = val
		}
		if val > max {
			max = val
		}
	}

	return max-min+1 <= 5
}

func main() {
	log.Println(IsContinuous([]int{0, 3, 2, 6, 4}))
}
