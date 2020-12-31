package main

import (
	"fmt"
)

func main() {
	//a := 0
	//fmt.Scan(&a)
	//fmt.Printf("%d\n", a)
	fmt.Println(compare("123", "321"))
}

func compare(s1, s2 string) bool {

	s1R := []rune(s1)
	s2R := []rune(s2)

	if len(s1R) != len(s2R) {
		return false
	}

	m := make(map[int32]int, 0)
	for _, i := range s1R {
		m[i]++
	}

	for _, i := range s2R {
		m[i]--
	}

	for _, val := range m {
		if val != 0 {
			return false
		}
	}

	return true
}
