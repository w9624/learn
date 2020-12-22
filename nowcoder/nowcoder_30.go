package main

import "fmt"

// 连续子数组最大和
// 动态规划 O(n)
// 状态转移方程 dp[i] = max(dp[i-1]+num[i], num[i])
func FindGreatestSumOfSubArray(array []int) int {

	if len(array) == 0 {
		return 0
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	result := 0
	dp := make([]int, len(array)+1)
	for i := 1; i < len(array); i++ {
		dp[i] = max(dp[i-1]+array[i], array[i])
		result = max(dp[i], result)
	}

	return result
}

func main() {
	fmt.Println(FindGreatestSumOfSubArray([]int{1, -2, 3, 10, -4, 7, 2, -5}))
}
