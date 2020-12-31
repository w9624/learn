package main

import (
	"fmt"
	"log"
)

func Permutation(str string) []string {
	if len(str) == 0 {
		return nil
	}

	chars := []rune(str)
	//sort.Sort(sortRunes(chars))

	//var reuslt []string
	for {
		//result = append(reuslt, "1")
		fmt.Println(string(chars))
		from := len(chars) - 1
		for i := len(chars) - 1; i > 0 && chars[i-1] >= chars[i]; i-- {
			from--
			fmt.Println(from)
		}

		if from == 0 {
			break
		}

		change := from
		for i := from; i < len(chars)-1 && chars[i-1] <= chars[i]; i++ {
			change++
		}

		chars[from], chars[change] = chars[change], chars[from]

		reverse(chars, from+1, len(chars)-1)
	}

	return nil
}

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}
func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s sortRunes) Len() int {
	return len(s)
}

func main() {
	log.Printf("%v", Permutation("eacb"))

	//chars := []int{2, 3, 4, 5, 6, 7, 8, 9}
	//reverse(chars, 2, len(chars)-1)
	//fmt.Printf("%v", chars)
}

func reverse(arr []rune, from, to int) {
	for i := from; i < (from+to)/2; i++ {
		arr[i], arr[to-(i-from)] = arr[to-(i-from)], arr[i]
	}
}
