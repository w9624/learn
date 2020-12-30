package main

// 顺时针打印矩阵
// 技巧题：四次循环打印行和列
func printMatrix(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}

	row, cloumn := len(matrix)-1, len(matrix[0])-1
	top, bottom, left, right := 0, row, 0, cloumn

	result := make([]int, (row+1)*(cloumn+1))
	k := 0
	for top <= bottom && left <= right {
		for i := left; i <= right; i, k = i+1, k+1 {
			result[k] = matrix[top][i]
		}
		for i := top + 1; i <= bottom; i, k = i+1, k+1 {
			result[k] = matrix[i][right]
		}
		for i := right - 1; i >= left && top < bottom; i, k = i-1, k+1 {
			result[k] = matrix[bottom][i]
		}
		for i := bottom - 1; i > top && left < right; i, k = i-1, k+1 {
			result[k] = matrix[i][left]
		}
		top++
		bottom--
		left++
		right--
	}

	return result
}
