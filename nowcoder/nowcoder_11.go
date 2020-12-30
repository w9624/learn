package main

import "fmt"

// 二进制中1的个数
// 统计
// 1. 常规思路 map统计
// 2. 二进制 借位每次必定会少一个1，n&(n-1) 一定可以减少一个1
func NumberOf1(n int) int {

	result := 0

	// 每次必定减去一个1
	for n != 0 {
		result++
		n = n & (n - 1)
	}

	return result
}

func main() {
	fmt.Println(NumberOf1(10))
}
