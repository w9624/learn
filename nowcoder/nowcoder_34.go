package main

import "fmt"

func FirstNotRepeatingChar(str string) int {
	if len(str) == 0 {
		return -1
	}

	m := make(map[int32]int, 0)
	for _, s := range str {
		m[s]++
	}

	for i, s := range str {
		if m[s] == 1 {
			return i
		}
	}

	return -1
}

func main() {
	fmt.Println(FirstNotRepeatingChar("google"))
}
