package main

import (
	"git.ixiaochuan.cn/xclib/third_party/src/github.com/uber/tchannel-go/crossdock/log"
	"sort"
)

func Permutation(str string) []string {
	if len(str) == 0 {
		return nil
	}

	chars := []rune(str)
	//sort.Sort(sortRunes(chars))
	//sortStr := string(chars)
	ma := make(map[string]bool, 0)
	PermutationHelper(0, len(chars)-1, chars, ma)

	var ss []string
	for s, _ := range ma {
		ss = append(ss, s)
	}
	sort.Strings(ss)

	return ss
}

func PermutationHelper(m, n int, chars []rune, ma map[string]bool) {
	if m == n {
		s := string(chars)
		if _, ok := ma[s]; !ok {
			ma[s] = true
		}
	} else {
		swap := func(i, k int) {
			chars[i], chars[k] = chars[k], chars[i]
		}
		for i := m; i < n; i++ {
			swap(i, m)
			PermutationHelper(m+1, n, chars, ma)
			swap(i, m)
		}
	}
}

//type sortRunes []rune
//func (s sortRunes) Less(i, j int) bool {
//	return s[i] < s[j]
//}
//func (s sortRunes) Swap(i, j int) {
//	s[i], s[j] = s[j], s[i]
//}
//func (s sortRunes) Len() int {
//	return len(s)
//}

func main() {
	log.Printf("%v", Permutation("eacb"))
}
