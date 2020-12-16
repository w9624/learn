package main

import (
	"log"
	"strings"
)

// 替换空格
func ReplaceSpace(str string) string {

	var sb strings.Builder
	for _, s := range str {
		if s == ' ' {
			sb.WriteString("%20")
		} else {
			sb.WriteRune(s)
		}
	}

	return sb.String()
}

func main() {
	cases := []struct {
		str      string
		expected string
	}{
		{
			str:      "We Are Happy",
			expected: "We%20Are%20Happy",
		},
	}

	for _, c := range cases {
		if str := ReplaceSpace(c.str); c.expected != str {
			log.Fatalf("fail, source: %v, target: %v, expected: %v", c.str, str, c.expected)
		}
	}
}
