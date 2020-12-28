package main

import "log"

// 翻转单词顺序列
// 数组/串 -> 旋转
func ReverseSentence(ReverseSentence string) string {
	if len(ReverseSentence) <= 1 {
		return ReverseSentence
	}

	arr := []rune(ReverseSentence)
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[i]
	}

	return string(arr)
}

func main() {
	log.Println(ReverseSentence("abcXYZdef"))
}
