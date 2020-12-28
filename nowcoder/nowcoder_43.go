package main

import "log"

// 左旋转字符串
// 数组/串 -> 旋转
func LeftRotateString(str string, n int) string {
	if len(str) <= 1 {
		return str
	}

	if n > len(str) {
		n %= len(str)
	}

	arr := []rune(str)
	for i := 0; i < n/2; i++ {
		arr[i], arr[n-i-1] = arr[n-i-1], arr[i]
	}

	for i := 0; i < (len(arr)-n)/2; i++ {
		arr[i+n], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[i+n]
	}

	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[i]
	}

	return string(arr)
}

func main() {
	log.Println(LeftRotateString("abcXYZdef", 3))
}
