package main

import "fmt"

// 二进制中1的个数
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
